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

func main() {
	var jsonStr = `[{"age":22,"sex":"woman","username":"user3"},{"age":25,"sex":"man","username":"user4"}]`
	var user1 []User
	err := json.Unmarshal([]byte(jsonStr), &user1)
	if err != nil {
		fmt.Println("unmarshal failed, err:", err)
		return
	}

	fmt.Println(user1)

}
