package main

import "fmt"

func main() {
	s := "stringg"
	fmt.Println(removeduplicatecharacter(s))
}

func removeduplicatecharacter(s string) string {
	seen := make(map[rune]bool)
	var remove []rune

	for _, v := range s {
		if !seen[v] {
			seen[v] = true
			remove = append(remove, v)
		}
	}
	return string(remove)
}
