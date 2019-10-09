package main

import (
	"context"
	pt "demo/day18/grpc/my_grpc_proto"
	"fmt"
	"google.golang.org/grpc"
)

const (
	post = "127.0.0.1:18881"
)

func main() {
	conn, err := grpc.Dial(post, grpc.WithInsecure())
	if err != nil {
		fmt.Println("连接服务器失败", err)
	}
	defer conn.Close()

	//获得grpc句柄
	c := pt.NewHelloServerClient(conn)

	// 远程调用 SayHello接口
	rl, err := c.SayHello(context.Background(), &pt.HelloRequest{Name: "panda"})
	if err != nil {
		fmt.Println("cloud not get hello server ..", err)
		return
	}
	fmt.Println("hello server resp: ", rl.Message)

	//远程调用 GetHelloMsg接口
	msg, err := c.GetHelloMsg(context.Background(), &pt.HelloRequest{Name: "panda"})
	if err != nil {
		fmt.Println("cloud not get hello msg ..", err)
		return
	}

	fmt.Println("hello server resp: ", msg.Msg)

}
