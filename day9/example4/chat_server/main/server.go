package main

import (
	"fmt"
	"net"
)

func runServer(addr string) (err error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("listen failed :", err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed :", err)
			return err
		}

		go process(conn)

	}

}

func process(conn net.Conn) {
	defer conn.Close()
	client := &Client{
		conn: conn,
	}

	err := client.Process()
	if err != nil {
		fmt.Println("client process failed, ", err)
		return
	}

}
