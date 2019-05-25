package main

import (
	"demo/day9/chat/proto"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

func register(conn net.Conn, userId int, passwd string) (err error) {
	var msg proto.Message
	msg.Cmd = proto.UserRegister

	var registerCmd proto.RegisterCmd
	registerCmd.User.UserId = userId
	registerCmd.User.Passwd = passwd
	registerCmd.User.Nick = fmt.Sprintf("stu%d", userId)
	registerCmd.User.Sex = "man"

	data, err := json.Marshal(registerCmd)
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

	msg, err = readPackage(conn)
	if err != nil {
		fmt.Println("read package failed, err:", err)
		return
	}

	fmt.Println(msg)

	return
}
