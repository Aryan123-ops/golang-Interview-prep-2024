package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}

	for i := 1; i <= 3; i++ {
		rotatedArr := rotateRight(arr, i)
		fmt.Printf("Rotate %d steps to the right: %v\n", i, rotatedArr)
	}
}

func rotateRight(arr []int, steps int) []int {
	n := len(arr)
	if n == 0 {
		return arr
	}
	steps = steps % n
	return append(arr[n-steps:], arr[:n-steps]...)
}
