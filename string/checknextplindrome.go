package main

import (
	"fmt"
	"strconv"
)

// isPalindrome checks if a number is a palindrome
func isPalindrome(n int) bool {
	s := strconv.Itoa(n)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

// nextPalindrome finds the next palindrome number after a given number
func nextPalindrome(n int) int {
	n++
	for !isPalindrome(n) {
		n++
	}
	return n
}

func main() {
	number := 123
	fmt.Printf("The next palindrome after %d is %d\n", number, nextPalindrome(number))
}
