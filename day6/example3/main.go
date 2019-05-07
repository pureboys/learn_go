package main

import "fmt"

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

type ReadWrite interface {
	Reader
	Writer
}

type File struct {
}

func (f *File) Write() {
	fmt.Println("write data")
}

func (f *File) Read() {
	fmt.Println("read data")
}

func Test(rw ReadWrite) {
	rw.Read()
	rw.Write()
}

func main() {
	var f File
	Test(&f)
}
