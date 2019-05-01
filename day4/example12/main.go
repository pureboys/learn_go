package main

import (
	"fmt"
	"sort"
)

func testMapSort() {
	a := make(map[int]int, 5)

	a[8] = 10
	a[3] = 20
	a[2] = 130
	a[1] = 104
	a[17] = 107
	a[89] = 123

	var keys []int
	for k := range a {
		// fmt.Println(k, v)
		keys = append(keys, k)
	}

	sort.Ints(keys)
	for _, v := range keys {
		fmt.Println(v, a[v])
	}

}

func trans() {
	a := make(map[string]int)
	b := make(map[int]string)

	a["abc"] = 101
	a["efg"] = 10

	for k, v := range a {
		b[v] = k
	}

	fmt.Println(b)

}

func main() {
	//testMapSort()
	trans()
}
