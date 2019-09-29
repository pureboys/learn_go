package main

import (
	middleware_demo "demo/day17/middleware-demo"
	"fmt"
)

func main() {
	request := middleware_demo.NewRequest()
	request.RegisterMiddleware(Logger, Recovery, func(request *middleware_demo.Request) {
		fmt.Println("this is my logic")
	})

	request.Next()
}

func Recovery(request *middleware_demo.Request) {
	defer func() {
		recover()
		fmt.Println("i can catch the panic...")
	}()
	request.Next()
}

func Logger(request *middleware_demo.Request) {
	fmt.Println("start logger ...")
	request.Next()
	fmt.Println("end logger ...")
}
