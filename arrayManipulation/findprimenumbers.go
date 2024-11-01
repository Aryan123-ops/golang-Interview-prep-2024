package main

import "fmt"

func main() {
	for num := 2; num <= 10; num++ {
		if findprime(num) {
			fmt.Println(num)
		}
	}
}

func findprime(num int) bool {
	for i := 2; i*i <= num; i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}
