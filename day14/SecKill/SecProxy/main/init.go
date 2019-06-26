package main

import (
	"context"
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
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
		logs.Error("init loadSecConf failed, err: %v", err)
		return
	}

	initSecProductWatcher()

	logs.Info("init sec success")
	return
}

func initSecProductWatcher() {
	go watchSecProductKey(secKillConf.etcdConf.etcdSecProductKey)
}

func watchSecProductKey(key string) {

	rch := etcdClient.Watch(context.Background(), key)
	var secProductInfo []SecProductInfoConf
	var getConfSuccess = true

	for wresp := range rch {
		for _, ev := range wresp.Events {
			if ev.Type == mvccpb.DELETE {
				if len(secProductInfo) > 0 {
					secProductInfo = []SecProductInfoConf{}
				}
				logs.Warn("key[%s] config deleted", key)
				continue
			}

			if ev.Type == mvccpb.PUT && string(ev.Kv.Key) == key {
				err := json.Unmarshal(ev.Kv.Value, &secProductInfo)
				if err != nil {
					logs.Error("key[%s], unmarshal[%s] err: %v", ev.Kv.Key, ev.Kv.Value, err)
					getConfSuccess = false
					continue
				}
			}

			logs.Debug("%s %q: %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}

		if getConfSuccess {
			logs.Debug("get config from etcd success, %v", secProductInfo)
			updateSecProductInfo(secProductInfo)
		}

	}
}

func updateSecProductInfo(secProductInfo []SecProductInfoConf) {

}

// 读取etcd中的货物参数
func loadSecConf() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(secKillConf.etcdConf.timeout)*time.Second)
	response, err := etcdClient.Get(ctx, secKillConf.etcdConf.etcdSecProductKey)
	cancel()
	if err != nil {
		logs.Error("get [%s] from etcd failed , err: %v", secKillConf.etcdConf.etcdSecProductKey, err)
		return
	}

	var secProductInfo []SecProductInfoConf

	for key, value := range response.Kvs {
		logs.Debug("key[%s] value[%v]", key, value)
		err2 := json.Unmarshal(value.Value, &secProductInfo)
		if err2 != nil {
			err = err2
			logs.Error("unmarshal sec product info failed, err: %v", err)
			return
		}

		logs.Debug("sec info conf is [%v]", secProductInfo)
	}

	secKillConf.SecProductInfo = secProductInfo
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
