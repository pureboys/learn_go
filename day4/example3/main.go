package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(factor(5))
	recursive(0)
}

func recursive(n int) {
	fmt.Println("hello")
	time.Sleep(time.Second)
	if n > 10 {
		return
	}
	recursive(n + 1)
}

func factor(n int) int {
	if n == 1 {
		return 1
	}
	return factor(n-1) * n
}
