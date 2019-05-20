package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("get err:", err)
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("get data err:", err)
		return
	}

	fmt.Println(string(data))

}
