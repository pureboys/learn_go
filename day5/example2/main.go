package main

import (
	"fmt"
	"math/rand"
)

type Student struct {
	Name  string
	Age   int
	Score float32
	next  *Student
}

func main() {
	var head = &Student{
		Name:  "oliver",
		Age:   18,
		Score: 100,
	}

	//insertTail(head)
	//trans(head)
	//fmt.Println()

	insertHead(&head)
	delNode(head)

	var newNode = &Student{
		Name:  "stu1000",
		Age:   1122,
		Score: 100,
	}

	addNode(head, newNode)
	trans(head)

}

func delNode(student *Student) {
	var prev = student
	for student != nil {
		if student.Name == "stu6" {
			prev.next = student.next
			break
		}
		prev = student
		student = student.next
	}
}

func addNode(student *Student, newNode *Student) {
	for student != nil {
		if student.Name == "stu9" {
			newNode.next = student.next
			student.next = newNode
			break
		}
		student = student.next
	}
}

func insertHead(head **Student) {
	for i := 0; i < 10; i++ {
		var stu = &Student{
			Name:  fmt.Sprintf("stu%d", i),
			Age:   rand.Intn(100),
			Score: rand.Float32() * 100,
		}

		stu.next = *head
		*head = stu
	}
}

func insertTail(head *Student) {
	var tail = head
	for i := 0; i < 10; i++ {
		var stu = &Student{
			Name:  fmt.Sprintf("stu%d", i),
			Age:   rand.Intn(100),
			Score: rand.Float32() * 100,
		}
		tail.next = stu
		//fmt.Printf("tail-> %p, head -> %p " , tail, head)
		tail = stu
	}
}

func trans(p *Student) {
	for p != nil {
		fmt.Println(*p)
		p = p.next
	}
}
