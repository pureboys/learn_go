package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/ping", testPing)
	_ = r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

func testPing(c *gin.Context) {

	var s string
	s1 := make(chan string, 5)
	s2 := make(chan string, 5)

	go go1(s1)
	go go2(s2)

	ticker := time.NewTicker(time.Second * 2)
	select {
	case s = <-s1:
	case s = <-s2:
	case <-ticker.C:
		fmt.Println("time out...")
	}
	ticker.Stop()

	c.JSON(200, gin.H{
		"message": s,
	})
}

func go1(s chan string) {
	num := rand.Int31n(5)
	time.Sleep(time.Second * time.Duration(num))

	s <- "i am go1"
}

func go2(s chan string) {
	num := rand.Int31n(5)
	time.Sleep(time.Second * time.Duration(num))

	s <- "i am go2"
}
