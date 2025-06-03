package main

import "fmt"

func main() {
	fmt.Println(Factorial(4))
}

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}
