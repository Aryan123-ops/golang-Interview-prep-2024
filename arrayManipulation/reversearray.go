package main

import "fmt"

func reversestringArray(arr []string) []string {

	left := 0
	right := len(arr) - 1

	for left < right {
		// Swap the elements at left and right indices
		arr[left], arr[right] = arr[right], arr[left]

		// Move the indices toward the center
		left++
		right--
	}
	return arr
}

func main() {
	// Define an array of strings
	strArray := []string{"abc", "banana", "cherry", "date", "pig"}
	fmt.Println("Original Array:", strArray)

	// Reverse the array
	reversestringArray(strArray)
	// Prstring the reversed array
	fmt.Println("Reversed Array:", strArray)
}

func arrayValue(arr []string) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
