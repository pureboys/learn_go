package main

import (
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int, 10)
	go write(intChan)
	go read(intChan)

	time.Sleep(time.Second * 10)

}

func read(ch chan int) {
	for {
		var b int
		b = <-ch
		fmt.Println(b)
	}
}

func write(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
}
