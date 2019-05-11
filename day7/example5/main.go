package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	inputFile := "/home/oliver/go/src/demo/day7/example5/products.txt"
	outputFile := "/home/oliver/go/src/demo/day7/example5/products_copy.txt"

	bytes, err := ioutil.ReadFile(inputFile)
	if err != nil {
		_, err = fmt.Fprintf(os.Stdout, "File Error:%s", err)
		return
	}

	fmt.Printf("%s\n", string(bytes))
	err = ioutil.WriteFile(outputFile, bytes, 0644)
	if err != nil {
		panic(err.Error())
	}
}
