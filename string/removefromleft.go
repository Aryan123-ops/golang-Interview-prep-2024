package main

import "fmt"

func main() {
	s := "test##str#ing"

	strRune := []rune(s)
	temp := []string{}

	for _, v := range strRune {
		if v == '#' {
			temp = temp[:len(temp)-1]
		} else {
			temp = append(temp, string(v))
		}
	}
	fmt.Println(temp)
}

// package main

// import "fmt"

// func main() {
// 	s := "test##str#ing"

// 	strRune := []rune(s)
// 	temp := []rune{}

// 	i := 0
// 	for i < len(strRune) {
// 		if strRune[i] == '#' {
// 			// Skip the character immediately following the #
// 			i -= 2
// 		} else {
// 			temp = append(temp, strRune[i])
// 			i++
// 		}
// 	}
// 	fmt.Println(string(temp))
// }
