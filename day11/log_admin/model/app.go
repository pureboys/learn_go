package model

import (
	"github.com/astaxie/beego/logs"
	"github.com/jmoiron/sqlx"
)

type AppInfo struct {
	AppId       int    `db:"app_id"`
	AppName     string `db:"app_name"`
	AppType     string `db:"app_type"`
	CreateTime  string `db:"create_time"`
	DevelopPath string `db:"develop_path"`
	IP          []string
}

var (
	Db *sqlx.DB
)

func InitDB(db *sqlx.DB) {
	Db = db
}

func GetAllAppInfo() (appList []AppInfo, err error) {
	err = Db.Select(&appList, "select app_id, app_name, app_type, create_time, develop_path from tbl_app_info")
	if err != nil {
		logs.Warn("Get All App Info failed, err: %v", err)
		return
	}
	return
}

func GetIPInfoById(appId int) (ipList []string, err error) {
	err = Db.Select(&ipList, "select ip from tbl_app_ip where app_id=?", appId)
	if err != nil {
		logs.Warn("Get IP info by id failed, err: %v", err)
		return
	}
	return
}

func GetIpInfoByName(appName string) (ipList []string, err error) {
	var appId []int

	err = Db.Select(&appId, "select app_id from tbl_app_info where app_name=?", appName)
	if err != nil || len(appId) == 0 {
		logs.Warn("select app_id failed, Db.exec error : %v", err)
		return
	}

	err = Db.Select(&ipList, "select ip from  tbl_app_ip where app_id=?", appId[0])
	if err != nil {
		logs.Warn("Get All App Info failed, err:%v", err)
		return
	}
	return
}

func CreateApp(info *AppInfo) (err error) {
	conn, err := Db.Begin()
	if err != nil {
		logs.Warn("Create App failed, Db.begin error:%v", err)
		return
	}

	defer func() {
		if err != nil {
			_ = conn.Rollback()
			return
		}

		_ = conn.Commit()
	}()

	result, err := conn.Exec("insert into tbl_app_info(app_name, app_type, develop_path) values (?, ?, ?)", info.AppName, info.AppType, info.DevelopPath)
	if err != nil {
		logs.Warn("create app failed, Db.exec error:%v", err)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		logs.Warn("create app failed, Db.lastInsertId error:%v", err)
		return
	}

	for _, ip := range info.IP {
		_, errors := conn.Exec("insert into tbl_app_ip(app_id, ip) values (?,?)", id, ip)
		if errors != nil {
			err = errors
			logs.Warn("create app failed, conn.exec ip error: %v", err)
			return
		}
	}

	return
}
