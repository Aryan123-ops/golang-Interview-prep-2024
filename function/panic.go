package main

import "fmt"

func main() {
	Division(4, 2)
	Division(2, 0)
}

func Division(num1, num2 int) {
	if num2 == 0 {
		panic("can not divide any number by 0")
	} else {
		result := num1 / num2
		fmt.Println("Result is:", result)
	}
}
