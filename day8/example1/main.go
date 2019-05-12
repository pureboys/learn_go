package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Open("/home/oliver/go/src/demo/day8/example1/test.log")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var lineArr []byte

	for {
		line, isPrefix, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}

		lineArr = append(lineArr, line...)

		if !isPrefix {
			fmt.Printf("data: %s\n", string(lineArr))
			lineArr = lineArr[0:0]
		}

	}

}
