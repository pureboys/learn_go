package main

import "fmt"

type Student struct {
	name  string
	age   int
	score int
}

type People struct {
	name string
	age  int
}

type Test interface {
	Print()
	Sleep()
}

func (p *Student) Print() {
	fmt.Println("name,", p.name)
	fmt.Println("age,", p.age)
	fmt.Println("score,", p.score)
}

func (p *Student) Sleep() {
	fmt.Println(p.name + " is sleep ...")
}

func (p *Student) Eat() {

}

func (p *People) Print() {
	fmt.Println("name,", p.name)
	fmt.Println("age,", p.age)
}

func (p *People) Sleep() {
	fmt.Println(p.name + " is sleep ...")
}

func main() {
	var t Test
	t = &Student{
		name:  "stu1",
		age:   20,
		score: 100,
	}
	t.Print()
	t.Sleep()

	fmt.Println()

	var people = &People{
		name: "people",
		age:  100,
	}
	people.Print()
	people.Sleep()

}
