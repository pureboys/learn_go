package main

import "fmt"

func main() {

	var a = 10

	switch a {
	case 0, 1, 2, 3, 4, 5:
		fmt.Println("a is equal 0")
		fallthrough
	case 10:
		fmt.Println("a is equal 10")
		fallthrough
	default:
		fmt.Println("a is equal default")
	}

}
