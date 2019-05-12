package main

import (
	"encoding/json"
	"io/ioutil"
)

type Student struct {
	Name string
	Sex  string
	Age  int
}

func (p *Student) Save() (err error) {
	bytes, err := json.Marshal(p)
	if err != nil {
		return
	}

	err = ioutil.WriteFile("./stu.dat", bytes, 0755)
	return
}

func (p *Student) Load() (err error) {
	bytes, err := ioutil.ReadFile("./stu.dat")
	if err != nil {
		return
	}

	err = json.Unmarshal(bytes, p)
	return
}
