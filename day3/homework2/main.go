package main

import "fmt"

func main() {
	var num int

	_, _ = fmt.Scanf("%d", &num)

	for i := 1; i <= num; i++ {
		numOfPerfect(i)
	}

}

func numOfPerfect(num int) {
	var sum int
	for i := 1; i < num; i++ {
		if num%i == 0 {
			sum += i
		}
	}
	if sum == num {
		println(num)
	}

}
