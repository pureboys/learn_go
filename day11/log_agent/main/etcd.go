package main

import (
	"context"
	"demo/day11/log_agent/tailf"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"go.etcd.io/etcd/clientv3"
	"strings"
	"time"
)

type EtcdClient struct {
	client *clientv3.Client
	keys   []string
}

var (
	etcdClient *EtcdClient
)

func initEtcd(addr string, key string) (collectConf []tailf.CollectConf, err error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		logs.Error("connect failed, err:", err)
		return
	}

	etcdClient = &EtcdClient{
		client: cli,
	}

	if strings.HasSuffix(key, "/") == false {
		key = key + "/"
	}

	for _, ip := range localIPArr {
		etcdKey := fmt.Sprintf("%s%s", key, ip)
		etcdClient.keys = append(etcdClient.keys, etcdKey)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		response, err := cli.Get(ctx, etcdKey)
		if err != nil {
			logs.Error("client get from etcd failed, err:%v", err)
			continue
		}
		cancel()
		logs.Debug("resp from etcd:%v", response.Kvs)

		for _, value := range response.Kvs {
			if string(value.Key) == etcdKey {
				err := json.Unmarshal(value.Value, &collectConf)
				if err != nil {
					logs.Error("unmarshal failed, err:%v", err)
					continue
				}

				logs.Debug("log config is %v", collectConf)

			}
		}

	}

	initEtcdWatcher()

	return

}

func initEtcdWatcher() {
	for _, key := range etcdClient.keys {
		go watchKey(key)
	}
}

func watchKey(key string) {
	rch := etcdClient.client.Watch(context.Background(), key)
	var collectConf []tailf.CollectConf
	var getConfSuccess = true

	for wresp := range rch {
		for _, ev := range wresp.Events {
			if ev.Type == mvccpb.DELETE {
				logs.Warn("key[%s] config deleted", key)
				continue
			}

			if ev.Type == mvccpb.PUT && string(ev.Kv.Key) == key {
				err := json.Unmarshal(ev.Kv.Value, &collectConf)
				if err != nil {
					logs.Error("key[%s], unmarshal[%s] err: %v", ev.Kv.Key, ev.Kv.Value, err)
					getConfSuccess = false
					continue
				}
			}

			logs.Debug("%s %q: %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}

		if getConfSuccess {
			_ = tailf.UpdateConfig(collectConf)
		}

	}
}
