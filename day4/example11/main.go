package main

import (
	"fmt"
)

func testMap() {
	var a map[string]string
	a = make(map[string]string, 10)
	a["abc"] = "efg"
	a["abc"] = "efg"
	a["abc1"] = "efg"

	var b = map[string]string{
		"abc": "value",
	}

	fmt.Println(a)
	fmt.Println(b)
}

func testMap2() {
	a := make(map[string]map[string]string, 10)
	a["key1"] = make(map[string]string)
	a["key1"]["key21"] = "abc"
	a["key1"]["key22"] = "abc"
	a["key1"]["key23"] = "abc"
	a["key1"]["key24"] = "abc"
	a["key1"]["key25"] = "abc"

	fmt.Println(a)

}

func testMap3() {
	a := make(map[string]map[string]string, 100)

	modify(a)

	fmt.Println(a)
}

func modify(a map[string]map[string]string) {
	_, ok := a["zhangsan"]
	if !ok {
		a["zhangsan"] = make(map[string]string)
	}

	a["zhangsan"]["passwd"] = "654321"
	a["zhangsan"]["nickname"] = "ol2"

	return
}

func testMap4() {
	a := make(map[string]map[string]string, 10)
	a["key1"] = make(map[string]string)
	a["key1"]["key21"] = "abc"
	a["key1"]["key22"] = "abc"
	a["key1"]["key23"] = "abc"
	a["key1"]["key24"] = "abc"
	a["key1"]["key25"] = "abc"

	delete(a["key1"], "key25")

	for k, v := range a {
		fmt.Println(k)
		for k1, v1 := range v {
			fmt.Println("\t", k1, v1)
		}
	}

	fmt.Println(len(a["key1"]))

}

func testMap5() {
	a := make([]map[int]int, 5)

	//a[0] = make(map[int]int)
	//a[0][10] = 100

	for k := range a {
		a[k] = make(map[int]int)
	}

	a[2][10] = 200
	fmt.Println(a)
}

func main() {
	//testMap()
	//testMap2()
	//testMap3()
	//testMap4()
	testMap5()
}
