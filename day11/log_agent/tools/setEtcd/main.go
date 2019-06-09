package main

import (
	"context"
	"demo/day11/log_agent/tailf"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

const (
	EtcdKey = "/oliver/backend/logagent/config/10.0.0.8"
)

func main() {
	SetLogConfEtcd()
}

func SetLogConfEtcd() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	fmt.Println("connect success")
	defer cli.Close()

	var logConfArr []tailf.CollectConf

	logConfArr = append(logConfArr, tailf.CollectConf{
		LogPath: "/home/oliver/go/src/demo/day11/log_agent/logs/access.log",
		Topic:   "nginx_topic",
	})

	logConfArr = append(logConfArr, tailf.CollectConf{
		LogPath: "/home/oliver/go/src/demo/day11/log_agent/logs/error2.log",
		Topic:   "nginx_error_topic",
	})

	data, err := json.Marshal(logConfArr)
	if err != nil {
		fmt.Println("json marshal failed:", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, EtcdKey, string(data))
	cancel()
	if err != nil {
		fmt.Println("put failed, err:", err)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, EtcdKey)
	cancel()

	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}

	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}
