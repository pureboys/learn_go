package main

import (
	_ "demo/day14/SecKill/SecProxy/router"
	"github.com/astaxie/beego"
)

func main() {
	_ = beego.LoadAppConfig("ini", "conf/app.ini")

	err := initConfig()
	if err != nil {
		panic(err)
		return
	}

	err = initSec()
	if err != nil {
		panic(err)
		return
	}
	beego.Run()
}
