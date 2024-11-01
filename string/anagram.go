package main

import (
	"fmt"
	"sort"
	"strings"
)

// Helper function to sort a string
func sortString(s string) string {
	s = strings.ToLower(s) // Convert to lowercase to ensure case-insensitivity
	r := strings.Split(s, "")
	sort.Strings(r)
	return strings.Join(r, "")
}

// Function to check if two strings are anagrams
func areAnagrams(s1, s2 string) bool {
	s1 = strings.ReplaceAll(s1, " ", "")
	s2 = strings.ReplaceAll(s2, " ", "")
	return sortString(s1) == sortString(s2)
}

func main() {
	str1 := "Listen"
	str2 := "Silent"

	if areAnagrams(str1, str2) {
		fmt.Printf("%s and %s are anagrams\n", str1, str2)
	} else {
		fmt.Printf("%s and %s are not anagrams\n", str1, str2)
	}
}
