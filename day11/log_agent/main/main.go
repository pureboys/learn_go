package main

import (
	"demo/day11/log_agent/tailf"
	"fmt"
	"github.com/astaxie/beego/logs"
)

func main() {
	filename := "/home/oliver/go/src/demo/day11/log_agent/conf/log_agent.ini"
	err := loadConfig("ini", filename)
	if err != nil {
		fmt.Printf("load conf failed, err %v\n", err)
		panic("load conf failed")
		return
	}

	err = initLogger()
	if err != nil {
		fmt.Printf("load logger failed, err %v\n", err)
		panic("load logger failed")
		return
	}

	logs.Debug("load conf success, config: %v", appConfig)

	err = tailf.InitTail(appConfig.collectConf)
	if err != nil {
		logs.Error("init tail failed, err: %v", err)
		return
	}

	logs.Debug("initialize all success")
	err = serverRun()
	if err != nil {
		logs.Error("serverRun failed, err: %v", err)
		return
	}

	logs.Info("program exited")

}
