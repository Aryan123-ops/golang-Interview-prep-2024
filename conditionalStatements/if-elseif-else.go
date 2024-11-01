package main

import "fmt"

func main() {
	num := 99

	if num <= 50 {
		fmt.Println("Number is smaller or equals to 50")
		return
	} else if num >= 50 && num <= 100 {
		fmt.Println("Number is in between 51 and 100 ")
		return
	} else {
		fmt.Println("Number is greagter than 100")
		return
	}
}
