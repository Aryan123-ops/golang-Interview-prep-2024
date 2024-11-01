package main

import "fmt"

// The fallthrough keyword is used, which means that after executing the code for case 3, the program will continue executing the code in the next case (case 2), regardless of whether v matches that case.
func main() {
	v := 15 % 4
	switch v {
	case 3:
		fmt.Println(100)
		fallthrough
	case 2:
		fmt.Println(42)
	case 1:
		fmt.Println(10)
		fallthrough
	default:
		fmt.Println("default")
	}
}
