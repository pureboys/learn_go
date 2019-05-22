package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var form = `<html><body><form action="#" method="post" name="bar">
			<input type="text" name="in" />
			<input type="text" name="in" />
			<input type="submit" value="Submit" />
			</form></body></html>`

func main() {

	http.HandleFunc("/test1", logPanics(SimpleServer))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("some err:", err)
	}
}

func SimpleServer(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "<h1>hello, world</h1>")
}

func logPanics(handle http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			x := recover()
			if x != nil {
				log.Printf("[%v] caught panic: %v", request.RemoteAddr, x)
			}
		}()
		handle(writer, request)
	}
}
