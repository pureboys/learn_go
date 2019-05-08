package main

import "fmt"

type LinkNode struct {
	next *LinkNode
	data interface{}
}

type Link struct {
	head *LinkNode
	tail *LinkNode
}

func (p *Link) InsertHead(data interface{}) {
	node := &LinkNode{
		data: data,
		next: nil,
	}

	if p.tail == nil && p.head == nil {
		p.tail = node
		p.head = node
		return
	}

	node.next = p.head
	p.head = node
}

func (p *Link) InsertTail(data interface{}) {
	node := &LinkNode{
		data: data,
		next: nil,
	}

	if p.tail == nil && p.head == nil {
		p.tail = node
		p.head = node
		return
	}

	p.tail.next = node
	p.tail = node
}

func (p *Link) Trans() {
	p1 := p.head
	for p1 != nil {
		fmt.Println(p1.data)
		p1 = p1.next
	}
}

func main() {
	var intLink Link

	for i := 0; i < 10; i++ {
		// intLink.InsertHead(i)
		intLink.InsertTail(i)
	}
	intLink.Trans()

}
