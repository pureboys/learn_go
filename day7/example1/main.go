package main

import (
	"fmt"
	"os"
)

type Student struct {
	Name  string
	Age   int
	Score float32
}

func main() {
	// fmt.Fprintf(os.Stdout, "do do do")
	// openFile()

	var str = "stu01 18 89.2"
	var stu Student

	_, _ = fmt.Sscanf(str, "%s %d %f", &stu.Name, &stu.Age, &stu.Score)
	fmt.Println(stu)
}

func openFile() {
	file, err := os.OpenFile("./test.log", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open file err:", err)
	}

	_, err = fmt.Fprintf(file, "do balance err\n")
	err = file.Close()
}
