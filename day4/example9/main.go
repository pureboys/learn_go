package main

import "fmt"

func main() {
	//testSlice()
	//testCopy()
	//testString()
	testModifyString()
}

func testModifyString() {
	s := "我hello world"
	s1 := []rune(s)

	s1[0] = '0'
	str := string(s1)

	fmt.Println(str)
}

func testString() {
	s := "hello world"
	s1 := s[0:5]
	s2 := s[6:]
	fmt.Println(s1)
	fmt.Println(s2)
}

func testCopy() {
	var a = []int{1, 2, 3, 4, 5, 6}
	b := make([]int, 1)

	copy(b, a)
	fmt.Println(b)
}

func testSlice() {
	var a = [...]int{1, 2, 3, 4, 5}
	s := a[1:]

	s[1] = 100
	fmt.Printf("s=%p a[1]=%p\n", s, &a[1])
	fmt.Println("before a:", a)

	s = append(s, 10)
	s = append(s, 10)
	s = append(s, 10)
	s = append(s, 10)
	s = append(s, 10)

	s[1] = 1000
	fmt.Println("after a:", a)
	fmt.Println(s)
	fmt.Printf("s=%p a[1]=%p\n", s, &a[1])

}
