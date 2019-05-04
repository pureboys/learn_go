package main

import (
	"fmt"
	"time"
)

type Car struct {
	name string
	age  int
}

type Train struct {
	Car
	int
	start time.Time
	age   int
}

func main() {
	var t Train
	t.name = "train"
	t.age = 100
	t.int = 200
	fmt.Println(t)

	t.Car.age = 200
	t.Car.name = "train2"
	fmt.Println(t)

}
