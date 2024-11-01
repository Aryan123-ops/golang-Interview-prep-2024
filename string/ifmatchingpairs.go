package main

import (
	"fmt"
)

func isValid(s string) bool {
	// Map for matching closing and opening brackets
	pairMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
		'>': '<',
	}

	// Stack to keep track of opening brackets
	var stack []rune

	for _, char := range s {
		// If it's a closing bracket
		if opening, found := pairMap[char]; found {
			// Check if the stack is empty or the top doesn't match
			if len(stack) == 0 || stack[len(stack)-1] != opening {
				return false
			}
			// Pop the opening bracket from the stack
			stack = stack[:len(stack)-1]
		} else {
			// Push opening bracket to the stack
			stack = append(stack, char)
		}
	}

	// If the stack is empty, all brackets were matched
	return len(stack) == 0
}

func main() {
	input := "{}[[(]><>]{()"
	fmt.Println("Is the string valid?:", isValid(input))
}
