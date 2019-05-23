package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	UserID   int    `db:user_id`
	UserName string `db:user_name`
	Sex      string `db:sex`
	Email    string `db:email`
}

type Place struct {
	Country string `db:country`
	City    string `db:city`
	TelCode int    `db:tel_code`
}

var DB *sqlx.DB

func init() {
	db, err := sqlx.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed", err)
		return
	}
	DB = db
}

func main() {
	result, err := DB.Exec("insert into person (user_name,sex,email) values (?,?,?)", "stu001", "man", "stu01@qq.com")
	if err != nil {
		fmt.Println("exec failed", err)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("exec failed", err)
		return
	}

	fmt.Println("insert success:", id)

}
