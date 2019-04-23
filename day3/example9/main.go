package main

import "fmt"

func main() {

	str := "hello,world中国"
	for index, val := range str {
		fmt.Printf("index[%d] val[%c] len[%d]\n", index, val, len([]byte(string(val))))
	}

}
