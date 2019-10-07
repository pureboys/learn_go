package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 建立网络链接
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:10010")
	if err != nil {
		fmt.Println("网络链接失败")
		return
	}

	var pd int
	err = client.Call("Panda.Getinfo", 10010, &pd)
	if err != nil {
		fmt.Println("call 失败")
		return
	}

	fmt.Println("最后得到的值为:", pd)

}
