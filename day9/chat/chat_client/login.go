package main

import (
	"demo/day9/chat/common"
	"demo/day9/chat/proto"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

func login(conn net.Conn, userId int, passwd string) (err error) {
	var msg proto.Message
	msg.Cmd = proto.UserLogin

	var loginCmd proto.LoginCmd
	loginCmd.Id = userId
	loginCmd.Passwd = passwd

	data, err := json.Marshal(loginCmd)
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

	//fmt.Println(msg)
	var loginResp proto.LoginCmdRes
	_ = json.Unmarshal([]byte(msg.Data), &loginResp)

	fmt.Println("online user lists:")
	for _, id := range loginResp.User {
		if id == userId {
			continue
		}
		// fmt.Printf("user %d\n", value)
		user := &common.User{
			UserId: id,
		}

		onlineUserMap[user.UserId] = user

	}

	return
}
