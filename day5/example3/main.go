package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score float32
	left  *Student
	right *Student
}

func main() {

	var root = &Student{
		Name:  "oliver",
		Age:   18,
		Score: 100,
	}

	var left1 = &Student{
		Name:  "oliver1",
		Age:   22,
		Score: 90,
	}

	var right1 = &Student{
		Name:  "oliver2",
		Age:   33,
		Score: 80,
	}

	root.left = left1
	root.right = right1

	var left2 = &Student{
		Name:  "oliver11",
		Age:   23,
		Score: 91,
	}
	left1.left = left2

	trans(root)

}

func trans(student *Student) {
	if student == nil {
		return
	}
	fmt.Println(student)
	trans(student.left)
	trans(student.right)
}
