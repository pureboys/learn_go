package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()

	split := strings.Split(string(data), "+")
	if len(split) != 2 {
		fmt.Println("please input a+b")
		return
	}

	space1 := strings.TrimSpace(split[0])
	space2 := strings.TrimSpace(split[1])
	fmt.Println(multi(space1, space2))

}

func multi(str1, str2 string) (result string) {

	if len(str1) == 0 && len(str2) == 0 {
		result = "0"
		return
	}

	var index1 = len(str1) - 1
	var index2 = len(str2) - 1
	var left int

	for index1 >= 0 && index2 >= 0 {
		c1 := str1[index1] - '0'
		c2 := str2[index2] - '0'

		sum := int(c1) + int(c2) + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}
		c3 := sum%10 + '0'
		result = fmt.Sprintf("%c%s", c3, result)
		index1--
		index2--
	}

	for index1 >= 0 {
		c1 := str1[index1] - '0'
		sum := int(c1) + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}
		c3 := sum%10 + '0'
		result = fmt.Sprintf("%c%s", c3, result)
		index1--
	}

	for index2 >= 0 {
		c2 := str2[index2] - '0'
		sum := int(c2) + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}
		c3 := sum%10 + '0'
		result = fmt.Sprintf("%c%s", c3, result)
		index2--
	}

	if left == 1 {
		result = fmt.Sprintf("1%s", result)
	}

	return
}
