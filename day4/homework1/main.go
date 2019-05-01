package main

import "fmt"

// 冒泡排序

func bsort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return
}

func main() {
	b := []int{5, 3, 1, 6, 85, 34, 78, 54}
	bsort(b)
	fmt.Println(b)

}
