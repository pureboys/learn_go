package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"
)

var url = []string{
	"http://www.baidu.com",
	"http://google.com",
	"http://taobao.com",
}

func main() {

	for _, v := range url {

		client := http.Client{
			Transport: &http.Transport{
				DialContext: func(ctx context.Context, network, addr string) (conn net.Conn, e error) {
					return net.DialTimeout(network, addr, time.Second)
				},
			},
		}

		resp, err := client.Head(v)
		if err != nil {
			fmt.Printf("head %s failed, err:%v\n", v, err)
			continue
		}

		fmt.Printf("head succ, status:%v\n", resp.Status)
	}

}
