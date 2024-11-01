package main

import "fmt"

// func main() {
// 	slice := make([]int, 5, 10)

// 	slice = append(slice, 1)
// 	fmt.Println(slice)
// 	fmt.Println(cap(slice))
// }

func main() {
	slice := make([]int, 3, 10)
	slice = append(slice, 1, 2, 3, 5, 5, 6, 7, 8)
	fmt.Println(slice)
	fmt.Println(cap(slice))
}
