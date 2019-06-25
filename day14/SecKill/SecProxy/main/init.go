package main

import (
	"github.com/astaxie/beego/logs"
	"github.com/gomodule/redigo/redis"
	"time"
)

func initSec() (err error) {
	err = initRedis()
	if err != nil {
		logs.Error("init redis failed, err: %v", err)
		return
	}

	err = initEtcd()
	if err != nil {
		logs.Error("init etcd failed, err: %v", err)
		return
	}

	logs.Info("init sec success")
	return
}

func initEtcd() (err error) {
	return
}

func initRedis() (err error) {

	_ = &redis.Pool{
		MaxIdle:     secKillConf.redisConf.redisMaxIdle,
		MaxActive:   secKillConf.redisConf.redisMaxActive,
		IdleTimeout: time.Duration(secKillConf.redisConf.redisIdleTimeout) * time.Second,
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", secKillConf.redisConf.redisAddr)
		},
	}

	return
}
