package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CharCount struct {
	ChCount    int
	NumCount   int
	SpaceCount int
	OtherCount int
}

func main() {
	file, err := os.Open("/home/oliver/go/src/demo/day7/example4/test.log")
	if err != nil {
		fmt.Println("read file err:", err)
		return
	}
	defer file.Close()

	var count CharCount

	reader := bufio.NewReader(file)
	for {
		readString, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Printf("read file failed, err:%v", err)
			return
		}

		runes := []rune(readString)
		for _, v := range runes {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}
	}

	fmt.Printf("char count %d\n", count.ChCount)
	fmt.Printf("space count %d\n", count.SpaceCount)
	fmt.Printf("num count %d\n", count.NumCount)
	fmt.Printf("other count %d\n", count.OtherCount)

}
