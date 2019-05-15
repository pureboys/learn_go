package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed :", err)
		return
	}
	defer conn.Close()

	_, err = conn.Do("MSet", "books", "abc", 100, "efg", 300)
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := redis.Ints(conn.Do("MGet", "abc", "efg"))
	if err != nil {
		fmt.Println("get abc,efg failed:", err)
		return
	}

	for _, v := range r {
		fmt.Println(v)
	}

}
