package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "trace_id", 123456)
	ctx = context.WithValue(ctx, "session", "asdfasfsafas232")
	process(ctx)
}

func process(ctx context.Context) {
	ret, ok := ctx.Value("trace_id").(int)
	if !ok {
		ret = 1234
	}
	fmt.Printf("ret is %d\n", ret)

	s, ok := ctx.Value("session").(string)
	if !ok {
		s = "123"
	}
	fmt.Printf("s is %s\n", s)

}
