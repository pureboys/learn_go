package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	_, _ = CopyFile("/home/oliver/go/src/demo/day7/example8/target.txt", "/home/oliver/go/src/demo/day7/example8/source.txt")
	fmt.Println("Copy Done!")
}

func CopyFile(dstName string, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
