package main

import "fmt"

func main() {

	for i := 101; i <= 200; i++ {
		var flag = false
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = true
				break
			}
		}
		if !flag {
			fmt.Printf("%d is prime\n", i)
		}
	}

}
