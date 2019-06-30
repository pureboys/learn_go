package main

import (
	"context"
	"demo/day14/SecKill/SecProxy/service"
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

	// 初始化service
	service.InitService(secKillConf)

	initSecProductWatcher()

	logs.Info("init sec success")
	return
}

func initSecProductWatcher() {
	go watchSecProductKey(secKillConf.EtcdConf.EtcdSecProductKey)
}

func watchSecProductKey(key string) {

	rch := etcdClient.Watch(context.Background(), key)
	var secProductInfo []service.SecProductInfoConf
	var getConfSuccess = true

	for wresp := range rch {
		for _, ev := range wresp.Events {
			if ev.Type == mvccpb.DELETE {
				if len(secProductInfo) > 0 {
					secProductInfo = []service.SecProductInfoConf{}
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

func updateSecProductInfo(secProductInfo []service.SecProductInfoConf) {

	var tmp = make(map[int]*service.SecProductInfoConf, 1024)
	for _, value := range secProductInfo {
		productInfo := value
		tmp[value.ProductId] = &productInfo
	}

	// 在需要改变的地方加锁保护 这是优化点！
	secKillConf.RwSecProductLock.Lock()
	secKillConf.SecProductInfoMap = tmp
	secKillConf.RwSecProductLock.Unlock()
}

// 读取etcd中的货物参数
func loadSecConf() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(secKillConf.EtcdConf.Timeout)*time.Second)
	response, err := etcdClient.Get(ctx, secKillConf.EtcdConf.EtcdSecProductKey)
	cancel()
	if err != nil {
		logs.Error("get [%s] from etcd failed , err: %v", secKillConf.EtcdConf.EtcdSecProductKey, err)
		return
	}

	var secProductInfo []service.SecProductInfoConf

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

	updateSecProductInfo(secProductInfo)

	return
}

func initLogs() (err error) {
	config := make(map[string]interface{})
	config["filename"] = secKillConf.LogPath
	config["level"] = convertLogLevel(secKillConf.LogLevel)

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
		Endpoints:   []string{secKillConf.EtcdConf.EtcdAddr},
		DialTimeout: time.Duration(secKillConf.EtcdConf.Timeout) * time.Second,
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
		MaxIdle:     secKillConf.RedisBlackConf.RedisMaxIdle,
		MaxActive:   secKillConf.RedisBlackConf.RedisMaxActive,
		IdleTimeout: time.Duration(secKillConf.RedisBlackConf.RedisIdleTimeout) * time.Second,
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", secKillConf.RedisBlackConf.RedisAddr)
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
