package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strings"
)

var (
	secKillConf = &SecSkillConf{}
)

type RedisConf struct {
	redisAddr        string
	redisMaxIdle     int
	redisMaxActive   int
	redisIdleTimeout int
}

type EtcdConf struct {
	etcdAddr          string
	timeout           int
	etcdSecKeyPrefix  string
	etcdSecProductKey string
}

type SecSkillConf struct {
	redisConf      RedisConf
	etcdConf       EtcdConf
	logPath        string
	logLevel       string
	SecProductInfo []SecProductInfoConf
}

type SecProductInfoConf struct {
	ProductId int
	StartTime int
	EndTime   int
	Status    int
	Total     int
	Left      int
}

func initConfig() (err error) {
	redisAddr := beego.AppConfig.String("redis_addr")
	etcdAddr := beego.AppConfig.String("etcd_addr")
	logs.Debug("redis config success, redis addr: %v, ", redisAddr)
	logs.Debug("etcd config success, etcd_ addr: %v, ", etcdAddr)

	secKillConf.etcdConf.etcdAddr = etcdAddr
	secKillConf.redisConf.redisAddr = redisAddr
	secKillConf.logPath = beego.AppConfig.String("log_path")
	secKillConf.logLevel = beego.AppConfig.String("log_level")

	if len(redisAddr) == 0 || len(etcdAddr) == 0 {
		err = fmt.Errorf("init config failed, redis[%s] or etcd[%s] config is null", redisAddr, etcdAddr)
		return
	}

	redisMaxIdle, err := beego.AppConfig.Int("redis_max_idle")
	if err != nil {
		err = fmt.Errorf("init config failed, read  redis_max_idle error: %v", err)
		return
	}

	redisMaxActive, err := beego.AppConfig.Int("redis_max_active")
	if err != nil {
		err = fmt.Errorf("init config failed, read  redis_max_active error: %v", err)
		return
	}

	redisIdleTimeout, err := beego.AppConfig.Int("redis_idle_timeout")
	if err != nil {
		err = fmt.Errorf("init config failed, read  redis_idle_timeout error: %v", err)
		return
	}

	etcdTimeout, err := beego.AppConfig.Int("etcd_timeout")
	if err != nil {
		err = fmt.Errorf("init config failed, read etcd_timeout error: %v", err)
		return
	}

	secKillConf.etcdConf.etcdSecKeyPrefix = beego.AppConfig.String("etcd_sec_key_prefix")
	if len(secKillConf.etcdConf.etcdSecKeyPrefix) == 0 {
		err = fmt.Errorf("init config failed, read etcd_sec_key error: %v", err)
		return
	}

	productKey := beego.AppConfig.String("etcd_product_key")
	if len(productKey) == 0 {
		err = fmt.Errorf("init config failed, read etcd_product_key error: %v", err)
		return
	}

	secKillConf.redisConf.redisMaxIdle = redisMaxIdle
	secKillConf.redisConf.redisMaxActive = redisMaxActive
	secKillConf.redisConf.redisIdleTimeout = redisIdleTimeout

	secKillConf.etcdConf.timeout = etcdTimeout
	if strings.HasSuffix(secKillConf.etcdConf.etcdSecKeyPrefix, "/") == false {
		secKillConf.etcdConf.etcdSecKeyPrefix = secKillConf.etcdConf.etcdSecKeyPrefix + "/"
	}
	secKillConf.etcdConf.etcdSecProductKey = fmt.Sprintf("%s/%s", secKillConf.etcdConf.etcdSecKeyPrefix, productKey)

	return
}
