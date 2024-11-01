package main

import "fmt"

func main() {
	isEvenOrOdd := checkEven(10)
	fmt.Println(isEvenOrOdd)
}

func checkEven(num int) bool {
	if num%2 == 0 {
		fmt.Println(num, "is even")
		return true
	} else {
		fmt.Println(num, "is odd")
		return false
	}
}
