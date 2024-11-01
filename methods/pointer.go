package main

import "fmt"

func main() {
	a := 1
	b := &a
	c := *b
	fmt.Println(c)
	fmt.Println(b)
}
