package main

import "fmt"

func main() {
	array := []int{5, 4, 3, 2, 1}
	fmt.Println(Bubble(array))
}

func Bubble(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}
