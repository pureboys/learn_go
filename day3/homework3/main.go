package main

import "fmt"

func main() {

	var str string

	_, _ = fmt.Scanf("%s", &str)

	palindromic(str)

}

func palindromic(str string) {

	t := []rune(str)
	length := len(t)

	for i := range t {
		if i == length/2 {
			break
		}
		last := length - i - 1
		if t[i] != t[last] {
			fmt.Printf("%s is not palindromic", str)
			return
		}
	}

	fmt.Printf("%s is palindromic", str)
}
