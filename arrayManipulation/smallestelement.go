package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 0, -1}
	fmt.Println(smallest(arr))
}

func smallest(arr []int) int {
	min := arr[0]

	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return min
}
