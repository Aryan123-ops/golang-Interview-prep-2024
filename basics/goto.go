package main

import "fmt"

func main() {
	i := 1

start:
	fmt.Println(i)
	i++

	if i <= 5 {
		goto start // jumps back to the start label
	}
}
