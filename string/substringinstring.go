package main

import (
	"fmt"
	"strings"
)

func main() {
	// Define the main string and the substring
	mainStr := "Hello, this is a sample string for testing."
	subStr := "sample"

	// Find the index of the substring in the main string
	index := strings.Index(mainStr, subStr)
	if index != -1 {
		fmt.Printf("The substring \"%s\" is found at index %d in the main string.\n", subStr, index)
	} else {
		fmt.Printf("The substring \"%s\" is not found in the main string.\n", subStr)
	}
}
