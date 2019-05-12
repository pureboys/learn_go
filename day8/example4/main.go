package main

import "fmt"

type student struct {
	name string
}

func main() {
	var intChan chan int
	intChan = make(chan int, 10)
	intChan <- 10

	var mapChan chan map[string]string
	mapChan = make(chan map[string]string, 10)
	m := make(map[string]string, 16)
	m["stu01"] = "001"
	m["stu02"] = "002"

	mapChan <- m

	var stuChan chan interface{}
	stuChan = make(chan interface{}, 10)
	stu := &student{name: "stu01"}
	stuChan <- stu

	var stu01 interface{}
	stu01 = <-stuChan
	fmt.Println(stu01)

	var stu02 *student
	stu02, ok := stu01.(*student)
	if !ok {
		fmt.Println("stu02 is not student")
		return
	}
	fmt.Println(stu02)

}
