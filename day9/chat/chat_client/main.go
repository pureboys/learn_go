package main

import (
	"demo/day9/chat/proto"
	"fmt"
	"net"
)

var userId int
var passwd string
var msgChan chan proto.UserRecvMessageReq

func init() {
	msgChan = make(chan proto.UserRecvMessageReq, 1000)
}

func main() {

	_, _ = fmt.Scanf("%d %s\n", &userId, &passwd)

	conn, err := net.Dial("tcp", "0.0.0.0:10000")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}

	// 登录部分
	err = login(conn, userId, passwd)
	if err != nil {
		fmt.Println("login failed, err:", err)
		return
	}

	// 注册部分
	//err = register(conn, userId, passwd)
	//if err != nil {
	//	fmt.Println("register failed, err:", err)
	//	return
	//}

	go processServerMessage(conn)

	for {
		logic(conn)
	}

}
