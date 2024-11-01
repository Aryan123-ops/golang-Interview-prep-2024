package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4}
	arr2 := []int{5, 6, 7, 8}

	fmt.Println("Array 1", arr1, "before swapping")
	fmt.Println("Array 2", arr2, "before swapping")

	arr2, arr1 = arr1, arr2

	fmt.Println("Array 2 ", arr2, "after swapping")
	fmt.Println("Array 1", arr1, "after swapping")
}
