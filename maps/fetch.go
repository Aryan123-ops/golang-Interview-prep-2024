// Program is used to print the keys and its value after sorting the map.
// Map is unsorted in golang

package main

import (
	"fmt"
	"sort"
)

func main() {
	emp := map[string]int{
		"C": 10,
		"B": 20,
		"A": 30,
	}

	// Extract and sort the keys
	var keys []string
	for key := range emp {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// Iterate through sorted keys and print values
	for _, key := range keys {
		value := emp[key]
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}

// package main

// import "fmt"

// type student struct {
// 	rollNumber int
// }

// func main() {
// 	m := map[student]string{
// 		student{1}: "a",
// 		student{2}: "b",
// 		student{3}: "c",
// 	}

// 	for key, v := range m {

// 		fmt.Printf("%v - %v\n", key, v)
// 	}

// }

// package main

// import (
// 	"fmt"
// 	"sort"
// )

// type student struct {
// 	rollNumber int
// }

// func main() {
// 	// Create a map of students
// 	m := map[student]string{
// 		{1}: "a",
// 		{2}: "b",
// 		{3}: "c",
// 		{5}: "e",
// 		{4}: "d",
// 	}

// 	// Collect all the keys (student structs)
// 	keys := make([]student, 0, len(m))
// 	for k := range m {
// 		keys = append(keys, k)
// 	}

// 	// Sort the keys based on rollNumber
// 	sort.Slice(keys, func(i, j int) bool {
// 		return keys[i].rollNumber < keys[j].rollNumber
// 	})

// 	// Iterate over the sorted keys and fetch values
// 	for _, key := range keys {
// 		fmt.Printf("%v - %v\n", key, m[key])
// 	}
// }
