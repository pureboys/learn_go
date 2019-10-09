package main

import (
	"context"
	pt "demo/day18/grpc/my_grpc_proto"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

const (
	host = "127.0.0.1:18881"
)

//对象要和proto内定义的服务一样
type server struct {
}

//实现RPC SayHello 接口
func (p *server) SayHello(ctx context.Context, im *pt.HelloRequest) (*pt.HelloReplay, error) {
	return &pt.HelloReplay{Message: "hello" + im.Name}, nil
}

//实现RPC GetHelloMsg 接口
func (p *server) GetHelloMsg(ctx context.Context, im *pt.HelloRequest) (*pt.HelloMessage, error) {
	return &pt.HelloMessage{Msg: "this is from server! HAHA!"}, nil
}

func main() {
	ln, err := net.Listen("tcp", host)
	if err != nil {
		fmt.Println("网络异常", err)
	}

	// 创建个grpc句柄
	srv := grpc.NewServer()
	//将server结构体注册到 grpc服务中
	pt.RegisterHelloServerServer(srv, &server{})

	//监听grpc服务
	err = srv.Serve(ln)
	if err != nil {
		fmt.Println("网络启动异常")
	}

}
