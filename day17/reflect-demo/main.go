package main

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"reflect"
)

type Person struct {
	UserId   int
	UserName string
	Sex      string
	Email    string
}

var (
	db *gorm.DB
)

func main() {
	db, _ = gorm.Open("mysql", "root:@/test?charset=utf8&parseTime=True&loc=Local")
	var result []*Person
	if err := Query(&result, "select * from person where user_name=?", "stu001"); err == nil {
		for i := 0; i < len(result); i++ {
			fmt.Println(result[i])
			fmt.Println(*result[i])
		}
	}
}

func Query(result interface{}, sql string, values ...interface{}) error {
	// type1是*[]*Person
	type1 := reflect.TypeOf(result)
	if type1.Kind() != reflect.Ptr {
		return errors.New("第一个参数必须是指针")
	}

	// type2是[]*Person
	type2 := type1.Elem() // 解指针后的类型
	if type2.Kind() != reflect.Slice {
		return errors.New("第一个参数必须指向切片")
	}

	// type3是*Person
	type3 := type2.Elem()
	if type3.Kind() != reflect.Ptr {
		return errors.New("切片元素必须是指针类型")
	}

	// 发起SQL查询
	rows, _ := db.Raw(sql, values...).Rows()
	for rows.Next() {
		//  type3.Elem()Person, elem是*Person
		elem := reflect.New(type3.Elem())
		// fmt.Println(elem.Type())

		// 传入*Person
		_ = db.ScanRows(rows, elem.Interface())
		// reflect.ValueOf(result).Elem()是[]*Person，Elem是*Person，newSlice是[]*Person
		newSlice := reflect.Append(reflect.ValueOf(result).Elem(), elem)
		// 扩容后的slice赋值给*result
		// reflect.ValueOf(result).Elem()是[]*Person
		fmt.Println(reflect.ValueOf(result).Elem().Type())
		reflect.ValueOf(result).Elem().Set(newSlice)
	}

	return nil
}
