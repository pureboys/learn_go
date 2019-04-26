package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var space, num, charters, others int

	var str string
	Scans(&str)

	runes := []rune(str)
	for _, v := range runes {
		switch {
		case v >= 'a' && v <= 'z':
			fallthrough
		case v >= 'A' && v <= 'Z':
			charters++
		case v == ' ':
			space++
		case v >= '0' && v <= '9':
			num++
		default:
			others++
		}
	}

	fmt.Printf("space is %d, num is %d, charters is %d, others is %d", space, num, charters, others)

}

func Scans(a *string) {
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	*a = string(data)
}
