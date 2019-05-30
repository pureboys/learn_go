package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
)

func main() {
	config := make(map[string]interface{})
	config["filename"] = "/home/oliver/go/src/demo/day11/logs/log_collect.log"
	config["level"] = logs.LevelDebug

	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("marshal failed, err:", err)
		return
	}

	_ = logs.SetLogger(logs.AdapterFile, string(configStr))
	logs.Debug("this is a test, my name is %s", "stu01")
	logs.Trace("this is a test, my name is %s", "stu02")
	logs.Warn("this is a test, my name is %s", "stu03")

}
