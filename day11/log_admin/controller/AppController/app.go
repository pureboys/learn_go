package AppController

import (
	"demo/day11/log_admin/model"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/prometheus/common/log"
	"strings"
)

type AppController struct {
	beego.Controller
}

func (p *AppController) AppList() {

	log.Debug("enter index controller")
	p.Layout = "layout/layout.html"

	appList, err := model.GetAllAppInfo()
	if err != nil {
		p.Data["Error"] = fmt.Sprintf("服务器繁忙")
		p.TplName = "app/error.html"
		logs.Warn("get app list failed, err: %v", err)
		return
	}

	log.Debug("app get list sucess, data: %v", appList)
	p.Data["applist"] = appList

	p.TplName = "app/index.html"
}

func (p *AppController) AppCreate() {

	appName := p.GetString("app_name")
	appType := p.GetString("app_type")
	developPath := p.GetString("develop_path")
	ipStr := p.GetString("iplist")

	p.Layout = "layout/layout.html"
	if len(appName) == 0 || len(appType) == 0 || len(developPath) == 0 || len(ipStr) == 0 {
		p.Data["Error"] = fmt.Sprintf("非法参数")
		p.TplName = "app/error.html"
		logs.Warn("invalid parameter")
		return
	}

	appInfo := &model.AppInfo{}
	appInfo.AppType = appType
	appInfo.AppName = appName
	appInfo.DevelopPath = developPath
	appInfo.IP = strings.Split(ipStr, ",")

	err := model.CreateApp(appInfo)
	if err != nil {
		p.Data["Error"] = fmt.Sprintf("创建项目失败，数据库繁忙")
		p.TplName = "app/error.html"
		logs.Warn("invalid parameter")
		return
	}

	p.Redirect("/app/list", 302)
}
