package main

import "fmt"

// 插入排序

func qsort(a []int, left, right int) {

	if left >= right {
		return
	}

	val := a[left]
	k := left
	for i := left + 1; i <= right; i++ {
		if a[i] < val {
			a[k] = a[i]
			a[i] = a[k+1]
			k++
		}
	}

	a[k] = val
	qsort(a, left, k-1)
	qsort(a, k+1, right)

}

func main() {
	b := []int{5, 3, 1, 6, 85, 34, 78, 54}
	qsort(b, 0, len(b)-1)
	fmt.Println(b)
}
