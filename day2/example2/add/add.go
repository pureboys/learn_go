package add

import (
	"fmt"
	_ "learn_go/day2/example2/test"
)

var Name = "hello world add"
var Age = 10

func init() {
	fmt.Println("init...")
	Name = "hello world add 2"
	Age = 100
}
