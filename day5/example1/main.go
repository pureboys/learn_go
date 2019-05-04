package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	score float32
}

func main() {

	var stu Student
	stu.Age = 18
	stu.Name = "oliver"
	stu.score = 80

	fmt.Println(stu)
	fmt.Printf("Name:%p ", &stu.Name)
	fmt.Printf("Age:%p ", &stu.Age)
	fmt.Println()

	var stu1 = Student{
		Age:  21,
		Name: "olvier2",
	}
	fmt.Println(stu1.Age)

}
