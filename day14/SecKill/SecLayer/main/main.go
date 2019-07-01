package main

import (
	_ "demo/day14/SecKill/SecProxy/router"
	"fmt"
	"github.com/astaxie/beego/logs"
)

func main() {

	err := initConfig("ini", "./conf/seclayer.ini")
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

	err = initSecKill()
	if err != nil {
		msg := fmt.Sprintf("init logger failed, err: %v", err)
		logs.Error(msg)
		panic(msg)
	}

	err = serviceRun()
	if err != nil {
		msg := fmt.Sprintf("service run return, err: %v", err)
		logs.Error(msg)
		return
	}

	logs.Info("service run exited")
}
