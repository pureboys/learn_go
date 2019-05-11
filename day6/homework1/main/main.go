package main

import (
	"demo/day6/homework1/balance"
	_ "demo/day6/homework1/thrid"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var insts []*balance.Instance
	for i := 0; i < 10; i++ {

		host := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255))
		port := 8080
		one := balance.NewInstance(host, port)
		insts = append(insts, one)
	}

	var banlanceName = "random"
	if len(os.Args) > 1 {
		banlanceName = os.Args[1]
	}

	for {
		inst, err := balance.DoBalance(banlanceName, insts)
		if err != nil {
			fmt.Println("do balance err:", err)
			continue
		}
		fmt.Println(inst)
		time.Sleep(time.Second)
	}

}
