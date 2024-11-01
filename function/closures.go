package main

import "fmt"

func main() {
	a := 5
	func() { // anonymous function designed
		fmt.Println("Value of a is ", a) // accesing value of a and this ability is called closures.
		return
	}()
}
