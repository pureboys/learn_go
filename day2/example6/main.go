package main

import "fmt"

func swap(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
	return
}

func main() {
	first := 100
	second := 200
	swap(&first, &second)
	fmt.Println("first", first)
	fmt.Println("second", second)
}
