package service

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"github.com/gomodule/redigo/redis"
	"math/rand"
	"time"
)

func initRedisPool(redisConf RedisConf) (pool *redis.Pool, err error) {
	pool = &redis.Pool{
		MaxIdle:     redisConf.RedisMaxIdle,
		MaxActive:   redisConf.RedisMaxActive,
		IdleTimeout: time.Duration(redisConf.RedisIdleTimeout) * time.Second,
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", redisConf.RedisAddr)
		},
	}
	conn := pool.Get()
	defer conn.Close()
	_, err = conn.Do("ping")
	if err != nil {
		logs.Error("ping redis failed")
		return
	}

	return
}

func initRedis(conf *SecLayerConf) (err error) {

	pool, err := initRedisPool(conf.Proxy2LayerRedis)
	if err != nil {
		logs.Error("init proxy2layer redis pool failed, err %v", err)
		return
	}
	secLayerContext.proxy2LayerRedisPool = pool

	pool2, err := initRedisPool(conf.Layer2ProxyRedis)
	if err != nil {
		logs.Error("init layer2proxy redis pool failed, err %v", err)
		return
	}
	secLayerContext.layer2ProxyRedisPool = pool2

	return
}

func RunProcess() (err error) {
	for i := 0; i < secLayerContext.secLayerConf.ReadGoroutineNum; i++ {
		secLayerContext.waitGroup.Add(1)
		go HandleReader()
	}

	for i := 0; i < secLayerContext.secLayerConf.WriteGoroutineNum; i++ {
		secLayerContext.waitGroup.Add(1)
		go HandleWrite()
	}

	for i := 0; i < secLayerContext.secLayerConf.HandleUserGoroutineNum; i++ {
		secLayerContext.waitGroup.Add(1)
		go HandleUser()
	}

	logs.Debug("all process go started")
	secLayerContext.waitGroup.Wait()
	logs.Debug("all process go exited")

	return

}

func HandleReader() {
	for {
		conn := secLayerContext.proxy2LayerRedisPool.Get()
		for {
			data, err := redis.String(conn.Do("BLPOP", "queuelist", 0))
			if err != nil {
				logs.Error("pop from queue failed, err : %v", err)
				break
			}

			logs.Debug("pop from queue, data: %s", data)

			var req SecRequest
			err = json.Unmarshal([]byte(data), &req)
			if err != nil {
				logs.Error("unmarshal data failed, err: %v", err)
				continue
			}

			// 如果超时则直接丢弃
			now := time.Now().Unix()
			if now-req.AccessTime.Unix() >= int64(secLayerContext.secLayerConf.MaxRequestWaitTimeout) {
				logs.Warn("req[%v] is expire", req)
				continue
			}

			timer := time.NewTicker(time.Millisecond * time.Duration(secLayerContext.secLayerConf.SendToHandleChanTimeout))
			select {
			case secLayerContext.Read2HandleChan <- &req:
			case <-timer.C:
				logs.Warn("send to handle chan timeout, req:%v", req)
				break
			}

			secLayerContext.Read2HandleChan <- &req
		}
		conn.Close()
	}
}

func HandleUser() {
	logs.Debug("handle user running")
	for req := range secLayerContext.Read2HandleChan {
		logs.Debug("begin process request:%v", req)
		res, err := HandleSecKill(req)
		if err != nil {
			logs.Warn("process request %v failed, err: %v", res, err)
			res = &SecResponse{
				Code: ErrServiceBusy,
			}
		}

		timer := time.NewTicker(time.Millisecond * time.Duration(secLayerContext.secLayerConf.SendToWriteChanTimeout))
		select {
		case secLayerContext.Handle2WriteChan <- res:
		case <-timer.C:
			logs.Warn("send to response chan timeout, res:%v", res)
			break
		}

	}
	return
}

func HandleSecKill(req *SecRequest) (res *SecResponse, err error) {
	secLayerContext.RwSecProductLock.RLock()
	defer secLayerContext.RwSecProductLock.RUnlock()

	res = &SecResponse{}
	product, ok := secLayerContext.secLayerConf.SecProductInfoMap[req.ProductId]
	if !ok {
		logs.Error("not found product:%v", req.ProductId)
		res.Code = ErrNotFoundProduct
		return
	}

	if product.Status == ProductStatusSoldout {
		res.Code = ErrSoldout
		return
	}

	now := time.Now().Unix()
	alreadySoldCount := product.secLimit.Check(now)
	if alreadySoldCount >= product.SoldMaxLimit {
		res.Code = ErrRetry
		return
	}

	secLayerContext.HistoryMapLock.Lock()
	userHistory, ok := secLayerContext.HistoryMap[req.UserId]
	if !ok {
		userHistory = &UserBuyHistory{
			history: make(map[int]int, 16),
		}

		secLayerContext.HistoryMap[req.UserId] = userHistory
	}

	historyCount := userHistory.GetProductBuyCount(req.ProductId)
	secLayerContext.HistoryMapLock.Unlock()

	if historyCount >= product.OnePersonBuyLimit {
		res.Code = ErrAlreadyBuy
		return
	}

	curSoldCount := secLayerContext.productCountMgr.Count(req.ProductId)
	if curSoldCount >= product.Total {
		res.Code = ErrSoldout
		product.Status = ProductStatusSoldout
		return
	}

	curRate := rand.Float64()
	if curRate > product.BuyRate {
		res.Code = ErrRetry
		return
	}

	userHistory.Add(req.ProductId, 1)
	secLayerContext.productCountMgr.Add(req.ProductId, 1)
	res.Code = ErrSecKillSucc

	return
}

func HandleWrite() {
	logs.Debug("handle write running")

	for res := range secLayerContext.Handle2WriteChan {
		err := sendToRedis(res)
		if err != nil {
			logs.Error("send to redis, err: %v, res:%v", err, res)
			continue
		}
	}

}

func sendToRedis(res *SecResponse) (err error) {
	data, err := json.Marshal(res)
	if err != nil {
		logs.Error("marshal failed, err: %v", err)
		return
	}
	conn := secLayerContext.layer2ProxyRedisPool.Get()
	_, err = redis.String(conn.Do("RPUSH", "layer2proxy_queue", string(data)))
	if err != nil {
		logs.Warn("rpush to redis failed, err:%v", err)
		return
	}
	return
}
