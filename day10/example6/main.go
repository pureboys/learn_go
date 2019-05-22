package main

import (
	"fmt"
	"html/template"
	"os"
)

type Person struct {
	Name string
	Age  string
}

func main() {

	files, err := template.ParseFiles("/home/oliver/go/src/demo/day10/example6/index.html")
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	p := Person{
		Name: "mary",
		Age:  "31",
	}

	err = files.Execute(os.Stdout, p)
	if err != nil {
		fmt.Println("There was an error:", err.Error())
	}

}
