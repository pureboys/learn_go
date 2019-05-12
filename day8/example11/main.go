package main

import (
	"fmt"
	"time"
)

func main() {
	//for {
	//	select {
	//	case <-time.After(time.Microsecond):
	//		fmt.Println("after")
	//	}
	//}

	for {
		t := time.NewTicker(time.Second)
		select {
		case <-t.C:
			fmt.Println("after")
		}
		t.Stop()
	}

}
