package main

import "fmt"

func test() {
	var a int8 = 100
	var b = int16(a)

	fmt.Printf("a=%d, b=%d", a, b)

}

func main() {
	test()
}
