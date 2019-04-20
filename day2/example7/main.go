package main

import "fmt"

func test() {
	var a = 100
	fmt.Println(a)

	for i := 0; i < 100; i++ {
		var b = i * 2
		fmt.Println(b)
	}
}

func main() {
	test()
}
