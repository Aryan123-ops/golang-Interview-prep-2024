package main

import "fmt"

func main() {
	str := "hello"
	substrings := findAllSubstrings(str)
	for _, substr := range substrings {
		fmt.Println(substr)
	}
}

func findAllSubstrings(s string) []string {
	var substrings []string
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			substrings = append(substrings, s[i:j])
		}
	}
	return substrings
}
