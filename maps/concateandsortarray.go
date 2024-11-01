package main

import (
	"fmt"
	"sort"
)

func main() {
	// Define two arrays
	arr1 := []int{1, 3, 5}
	arr2 := []int{2, 4, 6}

	// Concatenate the arrays
	concatenated := append(arr1, arr2...)

	// Create a map to hold the index-value pairs
	myMap := make(map[int]int)

	// Insert the concatenated array into the map
	for i, v := range concatenated {
		myMap[i] = v
	}

	// Fetch and print the values from the map in an ordered manner
	var keys []int
	for k := range myMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	fmt.Println("Ordered values from the map:")
	for _, k := range keys {
		fmt.Printf("Key: %d, Value: %d\n", k, myMap[k])
	}
}
