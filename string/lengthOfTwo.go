package main

import "fmt"

func main() {
	str := "hello"
	length := 2
	substrings := findSubstringsOfLength(str, length)
	for _, substr := range substrings {
		fmt.Println(substr)
	}
}

func findSubstringsOfLength(s string, length int) []string {
	var substrings []string
	for i := 0; i <= len(s)-length; i++ {
		substrings = append(substrings, s[i:i+length])
	}
	return substrings
}
