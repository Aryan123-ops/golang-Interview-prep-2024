package main

import (
	"fmt"
)

func main() {
	str1 := "hello"
	str2 := "hello"
	Comparison(str1, str2)
}

func Comparison(str1, str2 string) {
	if str1 == str2 {
		fmt.Printf("%s and %s are equal\n", str1, str2)
		return
	}
	fmt.Printf("%s and %s are not equal\n", str1, str2)
}
