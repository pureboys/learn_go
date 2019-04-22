package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "hello world abc  \n"

	result := strings.Replace(str, "world", "you", 1)
	fmt.Println(result)

	count := strings.Count(str, "l")
	fmt.Println(count)

	repeat := strings.Repeat(str, 3)
	fmt.Println(repeat)

	upper := strings.ToUpper(str)
	fmt.Println(upper)

	lower := strings.ToLower(str)
	fmt.Println(lower)

	space := strings.TrimSpace(str)
	fmt.Println(space)

	trim := strings.Trim(space, "\n\r")
	fmt.Println(trim)

	left := strings.TrimLeft(str, "\n\r")
	fmt.Println(left)

	right := strings.TrimRight(str, "\n\r")
	fmt.Println(right)

	fields := strings.Split(str, "l")
	for i := 0; i < len(fields); i++ {
		fmt.Println(fields[i])
	}

	join := strings.Join(fields, "l")
	fmt.Println(join)

	itoa := strconv.Itoa(1000)
	fmt.Println(itoa)

	number, err := strconv.Atoi(itoa)
	if err != nil {
		fmt.Println("can not convert to int ", err)
		return
	}

	fmt.Println(number)

}
