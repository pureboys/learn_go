package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type MyContext struct {
	context.Context
	Gin *gin.Context
}

type MyHandlerFunc func(c *MyContext)

func main() {
	r := gin.New()
	r.GET("/test", WithContext(func(myContent *MyContext) {
		dbQuery(myContent, "select * from abc")
		myContent.Gin.String(200, "完成")
	}))
	_ = r.Run()
}

func dbQuery(myContext *MyContext, sql string) {
	// 模拟长时间逻辑阻塞, 被context的5秒超时中断
	<-myContext.Done()
	// 模拟调用链埋点
	trace := myContext.Value("trace").(string)
	fmt.Println(trace)
}

func WithContext(my MyHandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		//fmt.Println(<-chan struct{}(nil))

		// 可以在gin.Context中设置key-value
		c.Set("trace", "假设这是一个调用链追踪sdk")

		//全局超时控制
		timeoutCtx, cancelFunc := context.WithTimeout(c, 10*time.Second)
		defer cancelFunc()

		myContext := MyContext{
			Context: timeoutCtx,
			Gin:     c,
		}
		my(&myContext)
	}

}
