package main

import (
	"fmt"
	"sort"
)

func testIntSort() {
	var a = [...]int{1, 8, 38, 25, 98, 2}
	sort.Ints(a[:])

	fmt.Println(a)
}

func testStrings() {
	var a = [...]string{"abc", "efg", "b", "A", "eee"}
	sort.Strings(a[:])

	fmt.Println(a)
}

func testFloat() {
	var a = [...]float64{2.3, 0.8, 28.9, 322, 6, 0.6}
	sort.Float64s(a[:])
	fmt.Println(a)
}

func testIntSearch() {
	var a = [...]int{1, 8, 38, 25, 98, 2}
	sort.Ints(a[:])
	index := sort.SearchInts(a[:], 25)

	fmt.Println(index)
}

func main() {
	// testIntSort()
	// testStrings()
	// testFloat()
	testIntSearch()
}
