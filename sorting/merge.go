package main

import (
	"fmt"
)

func main() {
	unsorted := []int{22, 12, 1, 23, 56}
	sorted := mergeSort(unsorted)
	fmt.Println(sorted)
}

func mergeSort(array []int) []int {
	if len(array) < 2 {
		return array
	}
	firsthalf := mergeSort(array[:len(array)/2])
	secondhalf := mergeSort(array[len(array)/2:])
	return merge(firsthalf, secondhalf)
}

func merge(a, b []int) []int {
	final := []int{}
	i := 0
	j := 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}
