package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var n int

	rand.Seed(time.Now().Unix())
	n = rand.Intn(100)

	for {
		var input int
		_, _ = fmt.Scanf("%d", &input)

		flag := false

		switch {
		case input == n:
			fmt.Println("you are right")
			flag = true
		case input > n:
			fmt.Println("bigger")
		case input < n:
			fmt.Println("less")
		}

		if flag {
			break
		}
	}

}
