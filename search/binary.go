package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5}
	target := 5
	result := binary(array, target)
	if result != -1 {
		fmt.Printf("Found %d at index %d\n", target, result)
	} else {
		fmt.Printf("Not found %d\n", target)
	}
}

func binary(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			low = mid - 1
		}
	}
	return -1
}
