package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/coreos/etcd/clientv3"
	"github.com/gomodule/redigo/redis"
	"time"
)

var (
	redisPool  *redis.Pool
	etcdClient *clientv3.Client
)

func initSec() (err error) {
	err = initLogs()
	if err != nil {
		logs.Error("init logs failed, err: %v", err)
		return
	}

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

	err = loadSecConf()
	if err != nil {
		return
	}

	logs.Info("init sec success")
	return
}

// 读取etcd中的货物参数
func loadSecConf() (err error) {
	key := fmt.Sprintf("%s/product", secKillConf.etcdConf.etcdSecKey)
	response, err := etcdClient.Get(context.Background(), key)
	if err != nil {
		logs.Error("get [%s] from etcd failed , err: %v", key, err)
		return
	}

	for key, value := range response.Kvs {
		logs.Debug("key[%s] value[%v]", key, value)
	}

	return
}

func initLogs() (err error) {
	config := make(map[string]interface{})
	config["filename"] = secKillConf.logPath
	config["level"] = convertLogLevel(secKillConf.logLevel)

	configStr, err := json.Marshal(config)
	if err != nil {
		logs.Error("marshal failed, err:", err)
		return
	}

	err = logs.SetLogger(logs.AdapterFile, string(configStr))
	if err != nil {
		logs.Error("SetLogger error, err:", err)
	}

	return
}

func initEtcd() (err error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{secKillConf.etcdConf.etcdAddr},
		DialTimeout: time.Duration(secKillConf.etcdConf.timeout) * time.Second,
	})

	if err != nil {
		logs.Error("connect failed, err:", err)
		return
	}

	etcdClient = cli
	return
}

func initRedis() (err error) {

	redisPool = &redis.Pool{
		MaxIdle:     secKillConf.redisConf.redisMaxIdle,
		MaxActive:   secKillConf.redisConf.redisMaxActive,
		IdleTimeout: time.Duration(secKillConf.redisConf.redisIdleTimeout) * time.Second,
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", secKillConf.redisConf.redisAddr)
		},
	}
	conn := redisPool.Get()
	defer conn.Close()
	_, err = conn.Do("ping")
	if err != nil {
		logs.Error("ping redis failed")
		return
	}
	return
}

func convertLogLevel(level string) int {
	switch level {
	case "debug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	case "info":
		return logs.LevelInfo
	case "trace":
		return logs.LevelTrace
	}
	return logs.LevelDebug
}
