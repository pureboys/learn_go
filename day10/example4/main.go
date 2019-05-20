package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/test1", SimpleServer)
	http.HandleFunc("/test2", FormServer)
	if err := http.ListenAndServe(":8080", nil); err != nil {

	}
}

var form = `<html><body><form action="#" method="post" name="bar">
			<input type="text" name="in" />
			<input type="text" name="in" />
			<input type="submit" value="Submit" />
			</form></body></html>`

func FormServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch r.Method {
	case "GET":
		io.WriteString(w, form)
	case "POST":
		r.ParseForm()
		io.WriteString(w, r.Form["in"][0])
		io.WriteString(w, "\n")
		io.WriteString(w, r.FormValue("in"))
	}
}

func SimpleServer(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>hello, world</h1>")
}
