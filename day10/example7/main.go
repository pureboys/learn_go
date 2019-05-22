package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Person struct {
	Name string
	Age  int
}

var myTemplate *template.Template

func main() {
	_ = iniTemplate("/home/oliver/go/src/demo/day10/example7/index.html")
	http.HandleFunc("/user/info", userInfo)
	err := http.ListenAndServe("0.0.0.0:8000", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}
}

func userInfo(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("handle hello")
	//fmt.Fprintf(writer, "hello ")

	person := Person{
		Name: "oliver",
		Age:  30,
	}

	myTemplate.Execute(writer, person)

	//file, err := os.OpenFile("/home/oliver/go/src/demo/day10/example7/test.log", os.O_CREATE|os.O_WRONLY, 0755)
	//if err != nil {
	//	fmt.Println("open failed err:", err)
	//	return
	//}
	//
	//myTemplate.Execute(file, person)
}

func iniTemplate(filename string) (err error) {
	myTemplate, err = template.ParseFiles(filename)
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	return
}
