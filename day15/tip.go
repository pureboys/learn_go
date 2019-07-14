package main

import (
	"html/template"
	"io/ioutil"
	"os"
)

type Page struct {
	Title string
	User  []User
}

type User struct {
	Username string
}

func main() {
	tpl := loadTemplate()

	data := Page{
		Title: "demo",
		User: []User{
			{Username: "oliver"},
			{Username: "hai"},
		},
	}
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

}

func loadTemplate() *template.Template {
	file, err := ioutil.ReadFile("tpl.gohtml")
	if err != nil {
		panic(err)
	}
	tpl, err := template.New("mytemplate").Parse(string(file))
	if err != nil {
		panic(err)
	}
	return tpl
}
