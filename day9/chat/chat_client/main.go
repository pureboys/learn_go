package main

import (
	"demo/day9/chat/proto"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"time"
)

func main() {
	var userId int
	var passwd string

	_, _ = fmt.Scanf("%d %s", &userId, &passwd)

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

	time.Sleep(time.Second * 100)

}

func readPackage(conn net.Conn) (msg proto.Message, err error) {
	var buf [8192]byte
	n, err := conn.Read(buf[0:4])
	if n != 4 {
		err = errors.New("read header failed")
		return
	}

	var packLen uint32
	packLen = binary.BigEndian.Uint32(buf[0:4])

	n, err = conn.Read(buf[0:packLen])
	if n != int(packLen) {
		err = errors.New("read body failed")
		return
	}

	fmt.Printf("receive data: %s\n", string(buf[0:packLen]))

	err = json.Unmarshal(buf[0:packLen], &msg)
	if err != nil {
		fmt.Println("unmarshal failed, err:", err)
		return
	}

	return
}

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
	for _, value := range loginResp.User {
		if value == userId {
			continue
		}
		fmt.Printf("user %d\n", value)
	}

	return
}

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
