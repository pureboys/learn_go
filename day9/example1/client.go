package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Error dialing ", err.Error())
		return
	}
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		trim := strings.Trim(input, "\r\n")
		if trim == "Q" {
			return
		}
		_, err := conn.Write([]byte(trim))
		if err != nil {
			return
		}
	}
}
