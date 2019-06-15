package AppController

import "github.com/astaxie/beego"

type AppController struct {
	beego.Controller
}

func (p *AppController) Index() {
	p.TplName = "app/index.html"

}
