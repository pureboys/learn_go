package LogController

import "github.com/astaxie/beego"

type LogController struct {
	beego.Controller
}

func (p *LogController) LogList() {

	p.Layout = "layout/layout.html"
	p.TplName = "log/"

}
