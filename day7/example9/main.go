package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Printf("len of args:%d\n", len(os.Args))
	for k, v := range os.Args {
		fmt.Printf("args[%d]=%s\n", k, v)
	}

}
