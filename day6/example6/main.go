package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a = 200
	test(a)

	var b = Student{
		Name:  "oliver",
		Age:   18,
		Score: 92,
	}
	test(b)

	var c = 1
	testInt(&c)
	fmt.Println(c)
}

type Student struct {
	Name  string
	Age   int
	Score float32
}

func test(b interface{}) {
	t := reflect.TypeOf(b)
	fmt.Println(t)

	v := reflect.ValueOf(b)
	fmt.Println(v.Kind())

	iv := v.Interface()
	stu, ok := iv.(Student)
	if ok {
		fmt.Printf("%v %T\n", stu, stu)
	}

}

func testInt(b interface{}) {
	val := reflect.ValueOf(b)
	val.Elem().SetInt(100)

	c := val.Elem().Int()
	fmt.Printf("get value interface{} %d \n", c)
	fmt.Printf("string val :%d\n", val.Elem().Int())
}
