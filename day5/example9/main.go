package main

import "fmt"

type Car struct {
	weight int
	name   string
}

type Bike struct {
	Car
	wheel int
}

type Train struct {
	c Car
}

func (p *Car) Run() {
	fmt.Println(p.name + " is running")
}

func (p *Train) String() string {
	str := fmt.Sprintf("name=[%s].weight=[%d]", p.c.name, p.c.weight)
	return str
}

func main() {

	var a Bike
	a.weight = 100
	a.name = "bike"
	a.wheel = 2

	fmt.Println(a)
	a.Run()

	var b Train
	b.c.weight = 100
	b.c.name = "train"
	b.c.Run()

	fmt.Printf("%s", &b)

}
