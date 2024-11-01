package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 3}
	ArrayAfterremove := removeduplicate(arr)
	fmt.Println(ArrayAfterremove)
}

func removeduplicate(arr []int) []int {
	seen := make(map[int]bool)
	removeArray := []int{}

	for _, num := range arr {
		if !seen[num] {
			seen[num] = true
			removeArray = append(removeArray, num)
		}
	}
	return removeArray
}
