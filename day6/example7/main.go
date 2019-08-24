package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string `json:"student_name"`
	Age   int
	Score float32
}

func main() {

	var a = Student{
		Name:  "stu-1",
		Age:   18,
		Score: 92.8,
	}

	TestStruck(&a)

	//result, _ := json.Marshal(a)
	//fmt.Println("json result:" , string(result))

}

func (s Student) Print() {
	fmt.Println("---------start-------")
	fmt.Println(s)
	fmt.Println("---------end-------")
}

func (s Student) Set(name string, age int, score float32) {
	s.Name = name
	s.Age = age
	s.Score = score
}

func TestStruck(a interface{}) {
	typeOf := reflect.TypeOf(a)
	fmt.Println(typeOf)

	val := reflect.ValueOf(a)
	fmt.Println(val.Type())

	kind := val.Kind()
	fmt.Println(kind)
	fmt.Println(val.Elem().Kind())

	fmt.Println(val.Elem())

	fmt.Println(val.Pointer())
	fmt.Println(val.Elem().UnsafeAddr())

	if kind != reflect.Ptr && val.Elem().Kind() == reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	num := val.Elem().NumField()
	val.Elem().Field(0).SetString("stu200")
	for i := 0; i < num; i++ {
		fmt.Printf("%d %v\n", i, val.Elem().Field(i).Kind())
	}

	fmt.Printf("struct has %d fields\n", num)

	tag := typeOf.Elem().Field(0).Tag.Get("json")
	fmt.Printf("tag=%s\n", tag)

	method := val.Elem().NumMethod()
	fmt.Printf("struct has %d method\n", method)

	val.Elem().Method(0).Call(nil)

}
