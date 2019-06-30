package service

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"time"
)

var (
	secKillConf *SecSkillConf
)

func InitService(serviceConf *SecSkillConf) {
	secKillConf = serviceConf
	logs.Debug("init service success, conf: %v", secKillConf)
}

func SecInfoById(productId int) (data map[string]interface{}, code int, err error) {

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
