package main

import (
	"fmt"
	"time"
)

func main() {
	go test()

	for {
		fmt.Println("i am running in main")
		time.Sleep(time.Second)
	}

}

func test() {
	var i int
	for {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}
