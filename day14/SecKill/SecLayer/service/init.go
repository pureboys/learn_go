package service

import (
	"github.com/astaxie/beego/logs"
	"github.com/coreos/etcd/clientv3"
	"time"
)

func InitSecLayer(conf *SecLayerConf) (err error) {

	err = initRedis(conf)
	if err != nil {
		logs.Error("init redis failed, err: %v", err)
		return
	}

	err = initEtcd(conf)
	if err != nil {
		logs.Error("init etcd failed, err: %v", err)
		return
	}

	err = loadProductFromEtcd(conf)
	if err != nil {
		logs.Error("init product failed, err: %v", err)
		return
	}

	return
}

func initEtcd(conf *SecLayerConf) (err error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{conf.EtcdConfig.EtcdAddr},
		DialTimeout: time.Duration(conf.EtcdConfig.Timeout) * time.Second,
	})

	if err != nil {
		logs.Error("connect failed, err:", err)
		return
	}

	secLayerContext.etcdClient = cli
	return
}
