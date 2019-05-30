package main

import (
	"fmt"
	"github.com/astaxie/beego/config"
)

func main() {
	conf, err := config.NewConfig("ini", "/home/oliver/go/src/demo/day11/config/log_collect.ini")
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return
	}

	port, err := conf.Int("server::port")
	if err != nil {
		fmt.Println("read server: port failed, err:", err)
		return
	}

	fmt.Println("Port:", port)
	logLevel, err := conf.Int("log::log_level")
	if err != nil {
		fmt.Println("read log_level failed, err:", err)
		return
	}

	fmt.Println("log_level:", logLevel)

	logPath := conf.String("log::log_path")
	fmt.Println("log_path:", logPath)

}
