package main

import "fmt"

// 选择排序

func ssort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		min := i
		for j := 1 + i; j < len(a); j++ {
			if a[min] > a[j] {
				min = j
			}
		}
		a[min], a[i] = a[i], a[min]
	}
	return
}

func main() {
	b := []int{5, 3, 1, 6, 85, 34, 78, 54}
	ssort(b)
	fmt.Println(b)
}
