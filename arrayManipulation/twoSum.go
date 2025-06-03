// return the index of the element in array that sums up to target .
package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	target := 5
	i, j := twosum(arr, target)
	fmt.Println(i, j) // prints the indices
}

func twosum(arr []int, target int) (int, int) {
	indexMap := make(map[int]int) // value -> index

	for i, num := range arr {
		complement := target - num
		if idx, ok := indexMap[complement]; ok {
			return idx, i
		}
		indexMap[num] = i
	}
	return -1, -1
}
