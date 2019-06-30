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
	redisBlackAddr := beego.AppConfig.String("redis_black_addr")
	etcdAddr := beego.AppConfig.String("etcd_addr")
	logs.Debug("redis config success, redis addr: %v, ", redisBlackAddr)
	logs.Debug("etcd config success, etcd_ addr: %v, ", etcdAddr)

	secKillConf.EtcdConf.EtcdAddr = etcdAddr
	secKillConf.RedisBlackConf.RedisAddr = redisBlackAddr
	secKillConf.LogPath = beego.AppConfig.String("log_path")
	secKillConf.LogLevel = beego.AppConfig.String("log_level")

	if len(redisBlackAddr) == 0 || len(etcdAddr) == 0 {
		err = fmt.Errorf("init config failed, redis[%s] or etcd[%s] config is null", redisBlackAddr, etcdAddr)
		return
	}

	redisMaxIdle, err := beego.AppConfig.Int("redis_black_max_idle")
	if err != nil {
		err = fmt.Errorf("init config failed, read  redis_max_idle error: %v", err)
		return
	}

	redisMaxActive, err := beego.AppConfig.Int("redis_black_max_active")
	if err != nil {
		err = fmt.Errorf("init config failed, read  redis_max_active error: %v", err)
		return
	}

	redisIdleTimeout, err := beego.AppConfig.Int("redis_black_idle_timeout")
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

	secKillConf.RedisBlackConf.RedisMaxIdle = redisMaxIdle
	secKillConf.RedisBlackConf.RedisMaxActive = redisMaxActive
	secKillConf.RedisBlackConf.RedisIdleTimeout = redisIdleTimeout

	secKillConf.EtcdConf.Timeout = etcdTimeout
	if strings.HasSuffix(secKillConf.EtcdConf.EtcdSecKeyPrefix, "/") == false {
		secKillConf.EtcdConf.EtcdSecKeyPrefix = secKillConf.EtcdConf.EtcdSecKeyPrefix + "/"
	}
	secKillConf.EtcdConf.EtcdSecProductKey = fmt.Sprintf("%s%s", secKillConf.EtcdConf.EtcdSecKeyPrefix, productKey)

	// 秒杀用户验证
	secKillConf.CookieSecretKey = beego.AppConfig.String("cookie_secret_key")

	// 用户秒杀每秒限速
	SecLimit, err := beego.AppConfig.Int("user_sec_access_limit")
	if err != nil {
		err = fmt.Errorf("init config failed, read user_sec_access_limit error: %v", err)
		return
	}
	secKillConf.UserSecAccessLimit = SecLimit

	// ip 秒杀每秒限速
	ipLimit, err := beego.AppConfig.Int("ip_sec_access_limit")
	if err != nil {
		err = fmt.Errorf("init config failed, read ip_sec_access_limit error: %v", err)
		return
	}
	secKillConf.IPSecAccessLimit = ipLimit

	referList := beego.AppConfig.String("refer_white_list")
	if len(referList) > 0 {
		secKillConf.ReferWhiteList = strings.Split(referList, ",")
	}

	return
}
