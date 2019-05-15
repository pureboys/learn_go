package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}
	defer conn.Close()
	_, err = conn.Do("expire", "abc", 10)
	if err != nil {
		fmt.Println(err)
		return
	}
}
