package main

import "fmt"

func main() {

	var str string

	_, _ = fmt.Scanf("%s", &str)

	palindromic(str)

}

func palindromic(str string) {
	lenStr := len(str)

	for i := 0; i <= lenStr/2; i++ {
		if str[i] != str[lenStr-1-i] {
			fmt.Printf("%s is not palindromic", str)
			return
		}
	}

	fmt.Printf("%s is palindromic", str)
}
