package main

import "fmt"

func feb(n int) int {
	if n <= 1 {
		return 1
	}
	return feb(n-1) + feb(n-2)
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(feb(i))
	}
}
