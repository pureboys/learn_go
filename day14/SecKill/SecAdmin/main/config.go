package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type MysqlConfig struct {
	UserName string
	Passwd   string
	Port     int
	Database string
	Host     string
}

type EtcdConf struct {
	Addr          string
	EtcdKeyPrefix string
	ProductKey    string
	Timeout       int
}

type Config struct {
	mysqlConf MysqlConfig
	etcdConf  EtcdConf
}

var (
	AppConf Config
)

func initConfig() (err error) {
	username := beego.AppConfig.String("mysql_user_name")
	if len(username) == 0 {
		logs.Error("load config of mysql_user_name failed, is null")
		return
	}

	AppConf.mysqlConf.UserName = username

	mysqlPasswd := beego.AppConfig.String("mysql_passwd")

	AppConf.mysqlConf.Passwd = mysqlPasswd

	mysqlHost := beego.AppConfig.String("mysql_host")
	if len(mysqlHost) == 0 {
		logs.Error("load config of mysql_host failed, is null")
		return
	}

	AppConf.mysqlConf.Host = mysqlHost

	mysqlDatabase := beego.AppConfig.String("mysql_database")
	if len(mysqlDatabase) == 0 {
		logs.Error("load config of mysql_database failed, is null")
		return
	}

	AppConf.mysqlConf.Database = mysqlDatabase

	port, err := beego.AppConfig.Int("mysql_port")
	if err != nil {
		logs.Error("load config of mysql_port failed, err:%v", err)
		return
	}

	AppConf.mysqlConf.Port = port

	addr := beego.AppConfig.String("etcd_addr")
	if len(addr) == 0 {
		logs.Error("load config of etcd_addr failed, is null")
		return
	}

	AppConf.etcdConf.Addr = addr

	prefix := beego.AppConfig.String("etcd_sec_key_prefix")
	if len(prefix) == 0 {
		logs.Error("load config of etcd_sec_key_prefix failed, is null")
		return
	}

	AppConf.etcdConf.EtcdKeyPrefix = prefix

	product := beego.AppConfig.String("etcd_product_key")
	if len(product) == 0 {
		logs.Error("load config of etcd_product_key failed, is null")
		return
	}

	AppConf.etcdConf.ProductKey = product

	timeout, err := beego.AppConfig.Int("etcd_timeout")
	if err != nil {
		logs.Error("load config of etcd_timeout failed, err:%v", err)
		return
	}

	AppConf.etcdConf.Timeout = timeout
	return
}
