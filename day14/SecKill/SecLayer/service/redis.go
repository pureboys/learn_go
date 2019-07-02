package service

import (
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
