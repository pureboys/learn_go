package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     16,
		MaxActive:   0,
		IdleTimeout: 300,
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", "0.0.0.0:6379")
		},
	}
}

func main() {

	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("Set", "abc", 100)
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := redis.Int(conn.Do("Get", "abc"))
	if err != nil {
		fmt.Println("get abc failed:", err)
		return
	}
	fmt.Println(r)

	pool.Close()

}
