package main

import "fmt"

func main() {
	arr := []int{0, 1, 0, 3, 5, 0, 12}
	sorted := Sort(arr)
	fmt.Println(sorted)

	fmt.Println(reverse(sorted))
}

func Sort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func reverse(arr []int) []int {
	left := 0
	right := len(arr) - 1

	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
	return arr
}
