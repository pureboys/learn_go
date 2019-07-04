package service

import (
	"context"
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"time"
)

func loadProductFromEtcd(conf *SecLayerConf) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	response, err := secLayerContext.etcdClient.Get(ctx, conf.EtcdConfig.EtcdSecProductKey)
	cancel()
	if err != nil {
		logs.Error("get [%s] from etcd failed , err: %v", conf.EtcdConfig.EtcdSecProductKey, err)
		return
	}

	var secProductInfo []SecProductInfoConf

	for key, value := range response.Kvs {
		logs.Debug("key[%s] value[%v]", key, value)
		err2 := json.Unmarshal(value.Value, &secProductInfo)
		if err2 != nil {
			err = err2
			logs.Error("unmarshal sec product info failed, err: %v", err)
			return
		}

		logs.Debug("sec info conf is [%v]", secProductInfo)
	}

	updateSecProductInfo(conf, secProductInfo)
	initSecProductWatcher(conf)

	return
}

func updateSecProductInfo(conf *SecLayerConf, secProductInfo []SecProductInfoConf) {

	var tmp = make(map[int]*SecProductInfoConf, 1024)
	for _, value := range secProductInfo {
		productInfo := value
		productInfo.secLimit = &SecLimit{}
		tmp[value.ProductId] = &productInfo
	}

	// 在需要改变的地方加锁保护 这是优化点！
	secLayerContext.RwSecProductLock.Lock()
	conf.SecProductInfoMap = tmp
	secLayerContext.RwSecProductLock.Unlock()
}

func initSecProductWatcher(conf *SecLayerConf) {
	go watchSecProductKey(conf)
}

func watchSecProductKey(conf *SecLayerConf) {
	key := conf.EtcdConfig.EtcdSecProductKey
	rch := secLayerContext.etcdClient.Watch(context.Background(), key)
	var secProductInfo []SecProductInfoConf
	var getConfSuccess = true

	for wresp := range rch {
		for _, ev := range wresp.Events {
			if ev.Type == mvccpb.DELETE {
				if len(secProductInfo) > 0 {
					secProductInfo = []SecProductInfoConf{}
				}
				logs.Warn("key[%s] config deleted", key)
				continue
			}

			if ev.Type == mvccpb.PUT && string(ev.Kv.Key) == key {
				err := json.Unmarshal(ev.Kv.Value, &secProductInfo)
				if err != nil {
					logs.Error("key[%s], unmarshal[%s] err: %v", ev.Kv.Key, ev.Kv.Value, err)
					getConfSuccess = false
					continue
				}
			}

			logs.Debug("%s %q: %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}

		if getConfSuccess {
			logs.Debug("get config from etcd success, %v", secProductInfo)
			updateSecProductInfo(conf, secProductInfo)
		}

	}
}
