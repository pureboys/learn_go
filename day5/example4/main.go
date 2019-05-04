package main

import "fmt"

type integer int

func main() {

	var i integer = 1000
	fmt.Println(i)

	j := int(i)
	fmt.Println(j)

}
