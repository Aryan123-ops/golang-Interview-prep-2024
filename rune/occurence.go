package main

import "fmt"

func main() {
	s := "hello"
	char := 'o'
	fmt.Println(countOccurence(s, char))
}

func countOccurence(s string, char rune) int {
	count := 0
	for _, c := range s {
		if c == char {
			count++
		}
	}
	return count
}
