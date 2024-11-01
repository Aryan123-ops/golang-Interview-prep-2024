package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	str := "hello world"
	vowels, consonant := countVowesAndConsonants(str)
	fmt.Printf("vowels are %d\n", vowels)
	fmt.Printf("consonant are %d\n", consonant)
}

func countVowesAndConsonants(s string) (int, int) {
	vowels := "aeiou"
	vowelsCount := 0
	consonantCount := 0

	s = strings.ToLower(s)
	for _, char := range s {
		if unicode.IsLetter(char) {
			if strings.ContainsRune(vowels, char) {
				vowelsCount++
			} else {
				consonantCount++
			}
		}
	}
	return vowelsCount, consonantCount
}
