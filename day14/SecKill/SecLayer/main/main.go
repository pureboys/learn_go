package main

import (
	"demo/day14/SecKill/SecLayer/service"
	_ "demo/day14/SecKill/SecProxy/router"
	"fmt"
	"github.com/astaxie/beego/logs"
)

func main() {

	err := initConfig("ini", "/home/oliver/go/src/demo/day14/SecKill/SecLayer/conf/seclayer.ini")
	if err != nil {
		msg := fmt.Sprintf("init config failed, err: %v", err)
		logs.Error(msg)
		panic(msg)
		return
	}

	err = initLogger()
	if err != nil {
		msg := fmt.Sprintf("init logger failed, err: %v", err)
		logs.Error(msg)
		panic(msg)
	}

	err = service.InitSecLayer(appConfig)
	if err != nil {
		msg := fmt.Sprintf("init logger failed, err: %v", err)
		logs.Error(msg)
		panic(msg)
	}

	logs.Debug("init sec layer success")

	err = service.Run()
	if err != nil {
		msg := fmt.Sprintf("service run return, err: %v", err)
		logs.Error(msg)
		return
	}

	logs.Info("service run exited")
}
