package service

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"github.com/gomodule/redigo/redis"
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

func HandleUser() {
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
			secLayerContext.Read2HandleChan <- &req
		}
		conn.Close()
	}
}

func HandleWrite() {

}

func HandleReader() {

}
