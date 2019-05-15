package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("start server...")
	listener, err := net.Listen("tcp", "0.0.0.0:50000")
	if err != nil {
		fmt.Println("listen failed , err:", err)
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			return
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		bytes := make([]byte, 512)
		_, err := conn.Read(bytes)
		if err != nil {
			fmt.Println("read err:", err)
			return
		}
		fmt.Println("read:", string(bytes))
	}

}
