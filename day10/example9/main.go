package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	UserID   int    `db:"user_id"`
	UserName string `db:"user_name"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"tel_code"`
}

var DB *sqlx.DB

func init() {
	db, err := sqlx.Open("mysql", "root:@tcp(127.0.0.1)/test")
	if err != nil {
		fmt.Println("open mysql failed", err)
		return
	}
	DB = db
}

func main() {
	var person []Person
	err := DB.Select(&person, "select user_id,user_name,sex,email from person where user_id = ?", 1)
	if err != nil {
		fmt.Println("exec failed:", err)
		return
	}

	fmt.Println("select success:", person)
}
