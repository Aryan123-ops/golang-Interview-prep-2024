// package main

// import "fmt"

// func main() {
// 	arr := []int{1, 1, 2, 2, 3, 4, 5, 5, 5}

// 	countMap := make(map[int]int)

// 	for _, num := range arr {
// 		countMap[num]++
// 	}

// 	for k, num := range countMap {
// 		fmt.Printf("%d:%d\n", k, num)
// 	}
// }

package main

import "fmt"

func main() {
	arr := "abbcdd"
	runes := []rune(arr)
	countMap := make(map[rune]int)
	for _, num := range runes {
		countMap[num]++
	}
	for k, v := range countMap {
		fmt.Printf("%c:%d\n", k, v)
	}
}
