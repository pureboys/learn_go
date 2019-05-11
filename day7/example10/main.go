package main

import (
	"flag"
	"fmt"
)

func main() {
	var confPath string
	var logLevel int
	flag.StringVar(&confPath, "c", "", "please input conf path")
	flag.IntVar(&logLevel, "d", 0, "please input log level")
	flag.Parse()

	fmt.Println("path:", confPath)
	fmt.Println("log Level:", logLevel)

}
