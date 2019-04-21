package main

import "fmt"

func main() {

	var n = 10
	var sum int

	var num = 1
	for j := 1; j <= n; j++ {
		num *= j
		fmt.Printf("%d!=%d\n", j, num)
		sum += num
	}

	fmt.Printf("sum is %d", sum)

}
