package main

import (
	"demo/day11/log_admin/model"
	_ "demo/day11/log_admin/router"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/jmoiron/sqlx"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	err := initDb()
	if err != nil {
		logs.Warn("initDb failed, err: %v", err)
		return
	}

	err = initEtcd()
	if err != nil {
		logs.Warn("initi etcd failed, err: %v", err)
		return
	}

	beego.Run()
}

func initEtcd() (err error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		logs.Warn("etcd connect failed err:%v", err)
		return
	}

	model.InitEtcd(cli)
	return
}

func initDb() (err error) {
	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/logadmin")
	if err != nil {
		logs.Warn("open mysql failed", err)
		return
	}

	model.InitDB(database)
	return
}
