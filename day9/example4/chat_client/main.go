package main

import (
	"demo/day9/example4/proto"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:10000")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}

	err = login(conn)
	if err != nil {
		fmt.Println("login failed, err:", err)
		return
	}

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

func login(conn net.Conn) (err error) {
	var msg proto.Message
	msg.Cmd = proto.UserLogin

	var loginCmd proto.LoginCmd
	loginCmd.Id = 1
	loginCmd.Passwd = "123456"

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

	fmt.Println(msg)

	return
}
