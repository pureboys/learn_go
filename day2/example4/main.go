package main

import (
	"fmt"
	"os"
)

func main() {
	getenv := os.Getenv("GOOS")
	fmt.Printf("The operating is: %s\n", getenv)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
}
