package main

import (
	"demo/day9/chat/proto"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func logic(conn net.Conn) {
	enterMenu(conn)
}

func enterMenu(conn net.Conn) {
	fmt.Println("1. list online user")
	fmt.Println("2. talk")
	fmt.Println("3. list message")
	fmt.Println("4. exit")

	var sel int
	_, _ = fmt.Scanf("%d\n", &sel)
	switch sel {
	case 1:
		outPutUserOnline()
	case 2:
		enterTalk(conn)
	case 3:
		listUnReadMsg()
	case 4:
		os.Exit(0)
	}

}

func listUnReadMsg() {
	select {
	case msg := <-msgChan:
		fmt.Println(msg.UserId, ":", msg.Data)
	default:
		return
	}
}

func enterTalk(conn net.Conn) {
	// var destUserId int
	var msg string
	fmt.Println("please input text")
	_, _ = fmt.Scanf("%s", &msg)
	_ = sendTextMessage(conn, msg)
}

func sendTextMessage(conn net.Conn, text string) (err error) {
	var msg proto.Message
	msg.Cmd = proto.UserSendMessageCmd

	var sendReq proto.UserSendMessageReq
	sendReq.Data = text
	sendReq.UserId = userId

	data, err := json.Marshal(sendReq)
	if err != nil {
		return
	}

	msg.Data = string(data)

	data, err = json.Marshal(msg)
	if err != nil {
		return
	}

	var buf [4]byte
	packLen := uint32(len(data))

	binary.BigEndian.PutUint32(buf[0:4], packLen)

	n, err := conn.Write(buf[:])
	if err != nil || n != 4 {
		fmt.Println("write dada failed")
		return
	}

	_, err = conn.Write([]byte(data))
	if err != nil {
		return
	}

	return
}
