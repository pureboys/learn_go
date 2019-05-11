package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	readString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("read string failed", err)
		return
	}

	fmt.Printf("read str succ, ret: %s\n", readString)

}
