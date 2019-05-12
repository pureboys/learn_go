package main

import "testing"

func TestStudent_Save(t *testing.T) {
	stu := &Student{
		Name: "stu01",
		Sex:  "man",
		Age:  10,
	}

	err := stu.Save()
	if err != nil {
		t.Fatalf("save student failed, err:%v", err)
	}

}

func TestStudent_Load(t *testing.T) {
	stu := &Student{}
	err := stu.Load()
	if err != nil {
		t.Fatalf("load student failed, err:%v", err)
	}

	if stu.Name != "stu01" {
		t.Fatalf("load student failed, name is wrong")
	}

	if stu.Sex != "man" {
		t.Fatalf("load student failed, sex is wrong")
	}

	if stu.Age != 10 {
		t.Fatalf("load student failed, age is wrong")
	}

}
