package main

import "fmt"

func main() {
	str := "racecar"
	if ispalindrome(str) {
		fmt.Println("is palindrome")
	} else {
		fmt.Println("is not palindrome")
	}
}

func ispalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] < s[len(s)-i-1] {
			return false
		}
	}
	return true
}
