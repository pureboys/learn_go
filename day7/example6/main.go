package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func main() {

	fName := "/home/oliver/go/src/demo/day7/example6/file.txt.gz"

	file, err := os.Open(fName)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v, Can't open %s: error: %s\n", os.Args[0], fName, err)
		os.Exit(1)
	}
	defer file.Close()

	fz, err := gzip.NewReader(file)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "open gzip failed, err: %v\n", err)
		return
	}

	reader := bufio.NewReader(fz)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("Done reading file")
			os.Exit(0)
		}
		fmt.Printf("%s", line)
		//fmt.Println(line)
	}

}
