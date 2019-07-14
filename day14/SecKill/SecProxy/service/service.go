package service

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego/logs"
	"time"
)

func NewSecRequest() (secRequest *SecRequest) {
	secRequest = &SecRequest{
		ResultChan: make(chan *SecResult, 1),
	}
	return
}

func SecInfoById(productId int) (data map[string]interface{}, code int, err error) {

	secKillConf.RwSecProductLock.RLock()
	defer secKillConf.RwSecProductLock.RUnlock()

	v, ok := secKillConf.SecProductInfoMap[productId]
	if !ok {
		code = ErrNotFoundProductId
		err = fmt.Errorf("not found product_id:%d", productId)
		return
	}

	start := false
	end := false
	status := "success"
	now := time.Now().Unix()

	if now-v.StartTime < 0 {
		start = false
		end = false
		status = "sec kill not start"
	}

	if now-v.StartTime > 0 {
		start = true
	}

	if now-v.EndTime > 0 {
		start = false
		end = true
		status = "sec kill is already end"
	}

	if v.Status == ProductStatusForceSaleOut || v.Status == ProductStatusSaleOut {
		start = false
		end = true
		status = "product is sale out"
	}

	data = make(map[string]interface{})
	data["product_id"] = productId
	data["start"] = start
	data["end"] = end
	data["status"] = status

	return
}

func SecInfo(productId int) (data []map[string]interface{}, code int, err error) {
	secKillConf.RwSecProductLock.RLock()
	defer secKillConf.RwSecProductLock.RUnlock()

	item, code, err := SecInfoById(productId)
	if err != nil {
		return
	}
	data = append(data, item)
	return
}

func SecInfoList() (data []map[string]interface{}, code int, err error) {
	secKillConf.RwSecProductLock.RLock()
	defer secKillConf.RwSecProductLock.RUnlock()

	for _, v := range secKillConf.SecProductInfoMap {

		item, _, err := SecInfoById(v.ProductId)
		if err != nil {
			logs.Error("get product_id [%d] failed, err: %v", v.ProductId, err)
			continue
		}

		data = append(data, item)
	}

	return
}

func userCheck(req *SecRequest) (err error) {

	found := false
	for _, refer := range secKillConf.ReferWhiteList {
		if refer == req.ClientReferer {
			found = true
			break
		}
	}

	if !found {
		err = fmt.Errorf("invalid request")
		logs.Warn("user[%d] is reject by refer, req[%v]", req.UserId, req)
		return
	}

	authData := fmt.Sprintf("%d:%s", req.UserId, secKillConf.CookieSecretKey)
	authSign := fmt.Sprintf("%x", md5.Sum([]byte(authData)))

	logs.Debug(authSign)
	logs.Debug(req.UserAuthSign)

	if authSign != req.UserAuthSign {
		err = fmt.Errorf("invalid user cookie auth")
		return
	}

	return
}

func SecKill(req *SecRequest) (data map[string]interface{}, code int, err error) {
	secKillConf.RwSecProductLock.RLock()
	defer secKillConf.RwSecProductLock.RUnlock()

	err = userCheck(req)
	if err != nil {
		code = ErrUserCheckAuthFailed
		logs.Warn("1. userId[%d] invalid, check failed, req[%v]", req.UserId, req)
		return
	}

	err = antiSpam(req)
	if err != nil {
		code = ErrUserServiceBusy
		logs.Warn("2. userId[%d] invalid, check failed, req[%v]", req.UserId, req)
		return
	}

	data, code, err = SecInfoById(req.ProductId)
	if err != nil {
		logs.Warn("3. userId[%d] secInfoBy Id failed, req[%v]", req.UserId, req)
		return
	}

	if code != 0 {
		logs.Warn("4. userId[%d] secInfoByid failed, code[%d] req[%v]", req.UserId, code, req)
		return
	}

	userKey := fmt.Sprintf("%d_%d", req.UserId, req.ProductId)
	secKillConf.UserConnMap[userKey] = req.ResultChan

	secKillConf.SecReqChan <- req

	ticker := time.NewTicker(time.Second * 10)

	defer func() {
		ticker.Stop()
		secKillConf.UserConnMapLock.Lock()
		delete(secKillConf.UserConnMap, userKey)
		secKillConf.UserConnMapLock.Unlock()
	}()

	select {
	case <-ticker.C:
		code = ErrProcessTimeout
		err = fmt.Errorf("request timeout")
		return
	case <-req.CloseNotify:
		code = ErrClientClosed
		err = fmt.Errorf("client already close")
		return
	case result := <-req.ResultChan:
		code = result.Code
		data["product_id"] = result.ProductId
		data["token"] = result.Token
		data["user_id"] = result.UserId
		return
	}

}
