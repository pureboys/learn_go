package main

import (
	"demo/day14/SecKill/SecLayer/service"
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/prometheus/common/log"
)

var (
	appConfig *service.SecLayerConf
)

func initConfig(confType, filename string) (err error) {
	conf, err := config.NewConfig(confType, filename)
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return
	}

	appConfig = &service.SecLayerConf{}
	appConfig.LogLevel = conf.String("logs::log_level")
	if len(appConfig.LogLevel) == 0 {
		appConfig.LogLevel = "debug"
	}

	appConfig.LogPath = conf.String("logs::log_path")
	if len(appConfig.LogPath) == 0 {
		appConfig.LogPath = "./logs"
	}

	appConfig.Proxy2LayerRedis.RedisAddr = conf.String("redis::redis_proxy2layer_addr")
	if len(appConfig.Proxy2LayerRedis.RedisAddr) == 0 {
		msg := "read redis::redis_proxy2layer_addr failed"
		log.Error(msg)
		err = fmt.Errorf(msg)
		return
	}

	appConfig.Proxy2LayerRedis.RedisMaxIdle, err = conf.Int("redis::redis_proxy2layer_idle")
	if err != nil {
		msg := "read redis::redis_proxy2layer_idle failed， err:%v"
		log.Error(msg, err)
		return
	}

	appConfig.Proxy2LayerRedis.RedisIdleTimeout, err = conf.Int("redis::redis_proxy2layer_idle_timeout")
	if err != nil {
		msg := "read redis::redis_proxy2layer_idle_timeout failed， err:%v"
		log.Error(msg, err)
		return
	}

	appConfig.Proxy2LayerRedis.RedisMaxActive, err = conf.Int("redis::redis_proxy2layer_active")
	if err != nil {
		msg := "read redis::redis_proxy2layer_active failed， err:%v"
		log.Error(msg, err)
		return
	}

	appConfig.WriteGoroutineNum, err = conf.Int("service::write_proxy2layer_goroutine_num")
	if err != nil {
		msg := "read service::write_proxy2layer_goroutine_num failed， err:%v"
		log.Error(msg, err)
		return
	}

	appConfig.ReadGoroutineNum, err = conf.Int("service::read_proxy2layer_goroutine_num")
	if err != nil {
		msg := "read service::read_proxy2layer_goroutine_num failed， err:%v"
		log.Error(msg, err)
		return
	}

	return
}
