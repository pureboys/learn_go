package main

import "fmt"

type slice struct {
	ptr *[5]int
	len int
	cap int
}

func make1(s slice, cap int) slice {
	s.ptr = new([5]int)
	s.cap = cap
	s.len = 0
	return s
}

func modify(s slice) {
	s.ptr[1] = 1000
}

func testSlice2() {
	var s1 slice
	s1 = make1(s1, 2)
	s1.ptr[0] = 100
	modify(s1)
	fmt.Println(s1.ptr)
}

func testSlice() {
	var slice []int
	var arr = [...]int{1, 2, 3, 4, 5}

	slice = arr[2:3]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice = slice[0:1]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

}

func testSlice3() {
	var b []int = []int{1, 2, 3, 4}
	modify1(b)
	fmt.Println(b)
}

func modify1(a []int) {
	a[1] = 1000
}

func testSlice4() {
	var a = [10]int{1, 2, 3, 4}
	b := a[1:5]
	fmt.Printf("%p\n", b)
	fmt.Printf("%p\n", &a[1])
}

func main() {
	testSlice()
	testSlice2()
	testSlice3()
	testSlice4()
}
