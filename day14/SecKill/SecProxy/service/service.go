package service

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/gomodule/redigo/redis"
	"strconv"
	"time"
)

var (
	secKillConf *SecSkillConf
)

func initBlackRedis() (err error) {

	secKillConf.blackRedisPool = &redis.Pool{
		MaxIdle:     secKillConf.RedisBlackConf.RedisMaxIdle,
		MaxActive:   secKillConf.RedisBlackConf.RedisMaxActive,
		IdleTimeout: time.Duration(secKillConf.RedisBlackConf.RedisIdleTimeout) * time.Second,
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", secKillConf.RedisBlackConf.RedisAddr)
		},
	}
	conn := secKillConf.blackRedisPool.Get()
	defer conn.Close()
	_, err = conn.Do("ping")
	if err != nil {
		logs.Error("ping redis failed")
		return
	}
	return
}

func InitService(serviceConf *SecSkillConf) {
	secKillConf = serviceConf
	loadBlackList()

	logs.Debug("init service success, conf: %v", secKillConf)
}

// 读取黑名单
func loadBlackList() (err error) {
	err = initBlackRedis()
	if err != nil {
		logs.Error("init black redis failed, err:%v", err)
		return
	}

	conn := secKillConf.blackRedisPool.Get()
	defer conn.Close()

	// id 黑名单
	reply, err := conn.Do("hgetall", "idblacklist")
	idList, err := redis.Strings(reply, err)
	if err != nil {
		logs.Warn("hget all failed, err: %v", err)
		return
	}

	for _, v := range idList {
		id, err2 := strconv.Atoi(v)
		err = err2
		if err != nil {
			logs.Warn("invalid user id [%v]", id)
			continue
		}
		secKillConf.idBlackMap[id] = true
	}

	// ip 黑名单
	reply, err = conn.Do("hgetall", "ipblacklist")
	ipList, err := redis.Strings(reply, err)
	if err != nil {
		logs.Warn("hget all failed, err: %v", err)
		return
	}

	for _, v := range ipList {
		secKillConf.ipBlackMap[v] = true
	}

	return
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
		logs.Warn("userId[%d] invalid, check failed, req[%v]", req.UserId, req)
		return
	}

	err = antiSpam(req)
	if err != nil {
		code = ErrUserServiceBusy
		logs.Warn("userId[%d] invalid, check failed, req[%v]", req.UserId, req)
		return
	}

	return
}
