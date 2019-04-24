package main

import "fmt"

var (
	result = func(a1 int, b1 int) int {
		return a1 + b1
	}
)

func test(a, b int) int {
	result := func(a1 int, b1 int) int {
		return a1 + b1
	}(a, b)

	return result
}

func main() {

	t := test(100, 200)
	println(t)

	println(result(300, 700))

	var i = 0
	defer fmt.Println(i)
	defer fmt.Println("second")

	i = 10
	fmt.Println(i)
}
