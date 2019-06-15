package router

import (
	"demo/day12/beego_example/controller/IndexController"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/index", &IndexController.IndexController{}, "*:Index")
}
