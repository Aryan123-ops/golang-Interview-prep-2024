package main

import (
	"fmt"
	"strings"
)

func main() {
	str := []string{"wel", "come", "to", "munger"}
	// str := "welocme"

	result := IfArrayOfStringContains(str)
	// result := ifSttringContains(str)
	fmt.Println(result)
}

// to check if [] strings contains the value
func IfArrayOfStringContains(s []string) bool {
	for _, v := range s {
		if strings.Contains(v, "wel") {
			return true
		} else {
			return false
		}
	}
	return false
}

// to check if string contains.
// func ifSttringContains(s string) bool {
// 	res := strings.Contains(s, "welo")
// 	return res
// }
