package main

import "fmt"

func main() {

	var n = 10
	var sum int

	for i := 1; i <= n; i++ {
		var num = 1
		for j := 1; j <= i; j++ {
			num *= j
		}
		sum += num
	}

	fmt.Printf("sum is %d", sum)

}
