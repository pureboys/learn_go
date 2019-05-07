package main

import "fmt"

type Stdudent struct {
	Name string
	Sex  string
}

func Test(a interface{}) {
	b, ok := a.(Stdudent)
	if ok == false {
		fmt.Println("convert failed")
		return
	}
	//b += 3
	fmt.Println(b)
}

func just(items ...interface{}) {
	for index, v := range items {
		switch v.(type) {
		case bool:
			fmt.Printf("%d params is bool. %v\n", index, v)
		case int, int32, int64:
			fmt.Printf("%d params is int. %v\n", index, v)
		case float32, float64:
			fmt.Printf("%d params is float. %v\n", index, v)
		case string:
			fmt.Printf("%d params is string. %v\n", index, v)
		case Stdudent:
			fmt.Printf("%d params is Student. %v\n", index, v)
		case *Stdudent:
			fmt.Printf("%d params is *Student. %v\n", index, v)
		}
	}
}

func main() {

	var b Stdudent
	Test(b)

	just(28, 8.3, "this is test", b, &b)

}
