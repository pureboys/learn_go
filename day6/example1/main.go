package main

import "fmt"

type Car interface {
	GetName() string
	Run()
	DiDi()
}

type Test interface {
	Hello()
}

type BMW struct {
	Name string
}

func (c *BMW) Hello() {
	fmt.Printf("hello, i am %s \n", c.Name)
}

func (c *BMW) GetName() string {
	return c.Name
}

func (c *BMW) Run() {
	fmt.Printf("%s is running\n", c.Name)
}

func (c *BMW) DiDi() {
	fmt.Printf("%s is didi\n", c.Name)
}

type BYD struct {
	Name string
}

func (c *BYD) GetName() string {
	return c.Name
}

func (c *BYD) Run() {
	fmt.Printf("%s is running\n", c.Name)
}

func (c *BYD) DiDi() {
	fmt.Printf("%s is didi\n", c.Name)
}

func main() {
	/*
		var a interface{}
		var b int
		var c float32

		a = b
		a = c

		fmt.Printf("type of a %T\n", a)
	*/

	var car Car
	var test Test

	bmw := &BMW{
		Name: "BMW",
	}
	car = bmw
	test = bmw
	car.Run()
	car.DiDi()
	test.Hello()

	byd := &BYD{
		Name: "BYD",
	}
	car = byd
	car.Run()
	car.DiDi()

}
