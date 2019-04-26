package main

import (
	"errors"
	"fmt"
	"time"
)

func initConfig() (err error) {
	return errors.New("init config failed")
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	err := initConfig()
	if err != nil {
		panic(err)
	}

	return
}

func main() {
	var i int
	fmt.Println(i)

	j := new(int)
	*j = 100
	fmt.Println(*j)

	var a []int
	a = append(a, 10, 20, 30)
	a = append(a, a...)
	fmt.Println(a)

	for {
		test()
		time.Sleep(time.Second)
	}

}
