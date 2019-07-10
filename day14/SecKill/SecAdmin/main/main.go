package main

import (
	_ "demo/day14/SecKill/SecAdmin/router"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	_ = beego.LoadAppConfig("ini", "conf/app.ini")
	err := initAll()
	if err != nil {
		panic(fmt.Sprintf("init database failed, err:%v", err))
		return
	}
	beego.Run()
}
