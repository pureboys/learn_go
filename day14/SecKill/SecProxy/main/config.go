package main

import (
	"demo/day14/SecKill/SecProxy/service"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strings"
)

var (
	secKillConf = &service.SecSkillConf{
		SecProductInfoMap: make(map[int]*service.SecProductInfoConf, 1024),
	}
)

func initConfig() (err error) {
	redisAddr := beego.AppConfig.String("redis_addr")
	etcdAddr := beego.AppConfig.String("etcd_addr")
	logs.Debug("redis config success, redis addr: %v, ", redisAddr)
	logs.Debug("etcd config success, etcd_ addr: %v, ", etcdAddr)

	secKillConf.EtcdConf.EtcdAddr = etcdAddr
	secKillConf.RedisConf.RedisAddr = redisAddr
	secKillConf.LogPath = beego.AppConfig.String("log_path")
	secKillConf.LogLevel = beego.AppConfig.String("log_level")

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

	secKillConf.EtcdConf.EtcdSecKeyPrefix = beego.AppConfig.String("etcd_sec_key_prefix")
	if len(secKillConf.EtcdConf.EtcdSecKeyPrefix) == 0 {
		err = fmt.Errorf("init config failed, read etcd_sec_key error: %v", err)
		return
	}

	productKey := beego.AppConfig.String("etcd_product_key")
	if len(productKey) == 0 {
		err = fmt.Errorf("init config failed, read etcd_product_key error: %v", err)
		return
	}

	secKillConf.RedisConf.RedisMaxIdle = redisMaxIdle
	secKillConf.RedisConf.RedisMaxActive = redisMaxActive
	secKillConf.RedisConf.RedisIdleTimeout = redisIdleTimeout

	secKillConf.EtcdConf.Timeout = etcdTimeout
	if strings.HasSuffix(secKillConf.EtcdConf.EtcdSecKeyPrefix, "/") == false {
		secKillConf.EtcdConf.EtcdSecKeyPrefix = secKillConf.EtcdConf.EtcdSecKeyPrefix + "/"
	}
	secKillConf.EtcdConf.EtcdSecProductKey = fmt.Sprintf("%s%s", secKillConf.EtcdConf.EtcdSecKeyPrefix, productKey)

	return
}
