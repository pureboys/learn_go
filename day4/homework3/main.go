package main

import "fmt"

// 插入排序

func isort(a []int) {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0; j-- {
			if a[j] > a[j-1] {
				break
			}
			a[j], a[j-1] = a[j-1], a[j]
		}
	}
	return
}

func main() {
	b := []int{5, 3, 1, 6, 85, 34, 78, 54}
	isort(b)
	fmt.Println(b)
}
