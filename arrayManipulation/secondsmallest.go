package main

import (
	"fmt"
)

func main() {
	array := []int{-1, -2, -3, 0, 1, 2, 3, 4}
	second := secondsmallest(array)
	fmt.Println(second)
}

func secondsmallest(arr []int) int {
	smallOne := 0
	smallTwo := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] < smallOne {
			smallTwo = smallOne
			smallOne = arr[i]
		} else if arr[i] < smallTwo && arr[i] != smallOne {
			smallTwo = arr[i]
		}
	}
	return smallTwo
}
