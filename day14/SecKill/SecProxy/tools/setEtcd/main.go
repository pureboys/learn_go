package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

type SecInfoConf struct {
	ProductId int
	StartTime int
	EndTime   int
	Status    int
	Total     int
	Left      int
}

const (
	EtcdKey = "/oliver/backend/secskill/product"
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

	var SecInfoConfArr []SecInfoConf

	SecInfoConfArr = append(SecInfoConfArr, SecInfoConf{
		ProductId: 1111,
		StartTime: 1561537530,
		EndTime:   1561824000,
		Status:    0,
		Total:     1000,
		Left:      1000,
	})

	SecInfoConfArr = append(SecInfoConfArr, SecInfoConf{
		ProductId: 2111,
		StartTime: 1561537530,
		EndTime:   1561824000,
		Status:    0,
		Total:     2000,
		Left:      2000,
	})

	data, err := json.Marshal(SecInfoConfArr)
	if err != nil {
		fmt.Println("json marshal failed:", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//
	//cli.Delete(ctx, EtcdKey)
	//return

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
