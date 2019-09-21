package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type MyContent struct {
	context.Context
	Gin *gin.Context
}

type MyHandleFunc func(c *MyContent)

func main() {
	r := gin.New()
	r.GET("/test", WithMyContent(func(c *MyContent) {
		dbQuery(c, "this is sql")
		c.Gin.String(http.StatusOK, "请求ok")
	}))
	_ = r.Run()
}

func dbQuery(ctx *MyContent, s string) {
	trace := ctx.Gin.Value("trace").(string)
	<-ctx.Done()
	fmt.Println(trace)
}

func WithMyContent(myContextHandle MyHandleFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 可以在gin.Context中设置key-value
		c.Set("trace", "假设这是一个调用链追踪sdk")

		// 全局超时控制
		timeoutCtx, cancelFunc := context.WithTimeout(c, 60*time.Second)
		defer cancelFunc()

		myCtx := MyContent{
			Context: timeoutCtx,
			Gin:     c,
		}

		myContextHandle(&myCtx)
	}
}
