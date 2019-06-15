package AppController

import "github.com/astaxie/beego"

type AppController struct {
	beego.Controller
}

func (p *AppController) Index() {

	p.Layout = "layout/layout.html"
	p.TplName = "app/index.html"

}
