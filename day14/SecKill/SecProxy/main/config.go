package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
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

type SecSkillConf struct {
	redisConf RedisConf
	etcdAddr  string
}

func initConfig() (err error) {
	redisAddr := beego.AppConfig.String("redis_addr")
	etcdAddr := beego.AppConfig.String("etcd_addr")
	logs.Debug("redis config success, redis addr: %v, ", redisAddr)
	logs.Debug("etcd config success, etcd_ addr: %v, ", etcdAddr)

	secKillConf.etcdAddr = etcdAddr
	secKillConf.redisConf.redisAddr = redisAddr

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

	secKillConf.redisConf.redisMaxIdle = redisMaxIdle
	secKillConf.redisConf.redisMaxActive = redisMaxActive
	secKillConf.redisConf.redisIdleTimeout = redisIdleTimeout

	return
}
