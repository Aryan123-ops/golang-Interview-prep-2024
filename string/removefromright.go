package main

import "fmt"

func main() {
	s := "test##st#r#in#g"

	strRune := []rune(s)
	temp := []rune{}

	i := 0
	for i < len(strRune) {
		if strRune[i] == '#' {
			// Skip the character immediately following the #
			i += 2
		} else {
			temp = append(temp, strRune[i])
			i++
		}
	}
	fmt.Println(string(temp))
}
