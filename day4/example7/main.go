package main

import "fmt"

func feb(n int) {
	var a []uint64
	a = make([]uint64, n)

	a[0] = 1
	a[1] = 1

	for i := 2; i < n; i++ {
		a[i] = a[i-1] + a[i-2]
	}

	for _, v := range a {
		fmt.Println(v)
	}

}

func testArray() {
	var a = [5]int{1, 2, 3, 4, 5}
	var a1 = [5]int{1, 2, 3, 4, 5}
	var a2 = [...]int{10, 11, 12, 33}
	var a3 = [...]int{1: 100, 2: 200}
	var a4 = [...]string{1: "hello", 3: "world"}

	fmt.Println(a)
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)

}
func testArray2() {
	var a = [...][5]int{
		{1, 2, 3, 4, 5},
		{11, 22, 33, 44, 55},
	}

	for k, v := range a {
		for k2, v2 := range v {
			fmt.Printf("(%d,%d)=%d ", k, k2, v2)
		}
		fmt.Println()
	}
}

func main() {
	feb(10)
	testArray()
	testArray2()
}
