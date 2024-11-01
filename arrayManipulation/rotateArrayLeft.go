package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	for i := 1; i <= 3; i++ {
		rotateArray := rotateLeft(arr, i)
		// fmt.Println(rotateArray)a
		fmt.Printf("Rotate %d steps to the right: %v\n", i, rotateArray)
	}
}

func rotateLeft(arr []int, steps int) []int {
	n := len(arr)

	if n == 0 {
		return arr
	}
	steps = steps % n
	return append(arr[steps:], arr[:steps]...)
}
