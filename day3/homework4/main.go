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

	for i := 0; i < len(str); i++ {
		fmt.Println(int(str[i]))
		switch {
		case int(str[i]) == 32:
			space++
		case int(str[i]) >= 48 && int(str[i]) <= 57:
			num++
		case (int(str[i]) >= 65 && int(str[i]) <= 90) || (int(str[i]) >= 97 && int(str[i]) <= 122):
			charters++
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
