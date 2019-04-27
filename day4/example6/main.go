package main

import "fmt"

func test1() {
	var a [10]int

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	for index, val := range a {
		fmt.Printf("a[%d]=%d\n", index, val)
	}
}

func test2() {
	var a [10]int
	b := a
	b[0] = 100
	fmt.Println(a)
}

func test3(arr *[5]int) {
	(*arr)[0] = 1000
}

func main() {
	// test1()
	// test2()
	var arr [5]int
	test3(&arr)
	fmt.Println(arr)
}
