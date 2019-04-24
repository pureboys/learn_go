package main

import "fmt"

type addFunc func(int, int) int

func add(a, b int) int {
	return a + b
}

func operator(op addFunc, a, b int) int {
	return op(a, b)
}

func main() {
	c := add
	sum := operator(c, 200, 300)
	fmt.Println(sum)
}
