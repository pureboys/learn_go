package main

import "fmt"

func main() {
	for i := 100; i <= 999; i++ {
		hundred := int(i / 100)
		ten := int((i % 100) / 10)
		unit := int((i % 100) % 10)

		if hundred*hundred*hundred+ten*ten*ten+unit*unit*unit == i {
			fmt.Printf("num is %d\n", i)
		}

	}
}
