package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/rpc"
)

type Panda int

func (this *Panda) Getinfo(argType int, replyType *int) error {
	fmt.Println("打印对端发送过来的内容为:", argType)

	// 修改内容值
	*replyType = argType + 10010

	return nil
}

func main() {
	// 页面请求
	http.HandleFunc("/panda", pandatext)

	// 实例化对象
	panda := new(Panda)
	// 注册一个服务对象
	_ = rpc.Register(panda)
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":10010")
	if err != nil {
		fmt.Println("net wrong")
		return
	}

	_ = http.Serve(listener, nil)

}

func pandatext(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world panda")
}
