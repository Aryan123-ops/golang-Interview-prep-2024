package main

import "fmt"

func main() {

	Division(4, 2)
	Division(4, 0)
	Division(8, 2)
}

func handlePanic() {
	a := recover()

	if a != nil {
		fmt.Println("recover", a)
	}
}

func Division(num1, num2 int) {
	defer handlePanic()
	if num2 == 0 {
		panic("Can not divide by 0")
	} else {
		result := num1 / num2
		fmt.Println("Result", result)
	}
}
