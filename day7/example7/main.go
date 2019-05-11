package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("/home/oliver/go/src/demo/day7/example7/output.data", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("An error occurred with file creation\n")
		return
	}

	defer file.Close()
	outPutWriter := bufio.NewWriter(file)
	outOutString := "hello world\n"
	for i := 0; i < 10; i++ {
		_, _ = outPutWriter.WriteString(outOutString)
	}
	_ = outPutWriter.Flush()

}
