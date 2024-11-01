package main

import "fmt"

func findDuplicates(arr []int) []int {
	seen := make(map[int]bool)
	duplicates := []int{}
	added := make(map[int]bool)

	for _, num := range arr {
		if seen[num] && !added[num] {
			duplicates = append(duplicates, num)
			added[num] = true
		} else {
			seen[num] = true
		}
	}

	return duplicates
}

func main() {
	numbers := []int{1, 2, 2, 2, 3, 4, 2, 5, 6, 3, 7, 8, 9, 1}

	duplicateValues := findDuplicates(numbers)

	if len(duplicateValues) > 0 {
		fmt.Println("Duplicate values in the array:", duplicateValues)
	} else {
		fmt.Println("No duplicate values found in the array.")
	}
}

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	array1 := [4]int{1, 2, 3, 4}
// 	array2 := [4]int{3, 4, 5, 6}

// 	// Concatenate arrays
// 	concatedArray := append(array1[:], array2[:]...)

// 	// Find duplicates
// 	duplicates := findDuplicates(concatedArray)

// 	// Print results
// 	fmt.Println("Concatenated Array:", concatedArray)
// 	fmt.Println("Duplicate Elements:", duplicates)
// }

// func findDuplicates(arr []int) []int {
// 	seen := make(map[int]bool)
// 	duplicates := []int{}

// 	for _, num := range arr {
// 		if seen[num] {
// 			// Found a duplicate
// 			duplicates = append(duplicates, num)
// 		} else {
// 			// First occurrence
// 			seen[num] = true
// 		}
// 	}

// 	return duplicates
// }
