package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("/home/oliver/go/src/demo/day7/example3/test.log")
	if err != nil {
		fmt.Println("read file err:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	readString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("read string failed, err:", err)
		return
	}

	fmt.Printf("read str succ, ret:%s\n", readString)

}
