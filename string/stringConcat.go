package main

import "fmt"

func main() {
	str1 := "Go"
	str2 := "is awesome"

	// result := fmt.Sprintf("%s %s", str1, str2)  using sprintf
	// fmt.Println(result)

	res := str1 + " " + str2
	fmt.Println(res)
}
