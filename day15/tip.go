package main

import (
	"demo/day15/server"
	"html/template"
	"io/ioutil"
	"net/http"
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

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := Page{
			Title: "demo",
			User: []User{
				{Username: "oliver"},
				{Username: "hai"},
			},
		}
		err := tpl.Execute(w, data)
		if err != nil {
			panic(err)
		}
	})

	srv := server.New(mux)

	err := srv.ListenAndServe()
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
