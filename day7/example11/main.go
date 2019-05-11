package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserName string
	NickName string
	Age      int
	Birthday string
	Sex      string
	Email    string
	Phone    string
}

func testStruct() {
	user1 := &User{
		UserName: "user1",
		NickName: "nickname1",
		Age:      18,
		Birthday: "2018/8/8",
		Sex:      "man",
		Email:    "abc@gmail.com",
		Phone:    "110",
	}

	data, err := json.Marshal(user1)
	if err != nil {
		fmt.Println("json marshal failed. err:", err)
		return
	}

	fmt.Printf("%s\n", string(data))
}

func testInt() {
	var age = 100
	data, err := json.Marshal(age)
	if err != nil {
		fmt.Println("json marshal failed. err:", err)
		return
	}

	fmt.Printf("%s\n", string(data))
}

func testMap() {
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["username"] = "user2"
	m["age"] = 20
	m["sex"] = "woman"

	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("json marshal failed. err:", err)
		return
	}

	fmt.Printf("%s\n", string(data))
}

func testSlice() {
	var s []map[string]interface{}
	m := make(map[string]interface{})
	m["username"] = "user3"
	m["age"] = 22
	m["sex"] = "woman"

	s = append(s, m)

	m = make(map[string]interface{})
	m["username"] = "user4"
	m["age"] = 25
	m["sex"] = "man"

	s = append(s, m)

	data, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json marshal failed. err:", err)
		return
	}

	fmt.Printf("%s\n", string(data))
}

func main() {

	// testStruct()
	// testInt()
	//testMap()

	testSlice()
}
