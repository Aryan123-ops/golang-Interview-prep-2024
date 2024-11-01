package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 45}
	fmt.Println(largest(arr))
}

func largest(arr []int) int {

	max := arr[0]

	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}
