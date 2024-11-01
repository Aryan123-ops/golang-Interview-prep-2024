package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3}
	arr2 := []int{4, 5, 6}
	fmt.Println("merge array", MergeTwoArrays(arr1, arr2))
}

func MergeTwoArrays(arr1, arr2 []int) []int {
	n1, n2 := len(arr1), len(arr2)
	mergeslice := make([]int, n1+n2)
	i, j, k := 0, 0, 0
	for i < n1 && j < n2 {
		if arr1[i] <= arr2[j] {
			mergeslice[k] = arr1[i]
			i++
		} else {
			mergeslice[k] = arr2[j]
			j++
		}
		k++
	}
	for i < n1 {
		mergeslice[k] = arr1[i]
		i++
		k++
	}
	for j < n2 {
		mergeslice[k] = arr2[j]
		j++
		k++
	}
	return mergeslice
}
