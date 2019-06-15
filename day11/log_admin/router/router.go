package router

import (
	"demo/day11/log_admin/controller/AppController"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/index", &AppController.AppController{}, "*:Index")
}
