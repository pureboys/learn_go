package main

import (
	"time"
)

func main() {

	for i := 0; i < 100; i ++ {
		go testGoroutine(i)
	}

	time.Sleep(time.Second)
}