// package main

// import "fmt"

// func main() {
// 	arr := []int{1, 1, 2, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 4}
// 	fmt.Println(mostRepeated(arr))
// }

// func mostRepeated(arr []int) int {
// 	counts := make(map[int]int)

// 	for _, num := range arr {
// 		counts[num]++
// 	}

// 	maxcount := 0
// 	mostrepeated := -1

// 	for num, count := range counts {
// 		if count > maxcount {
// 			maxcount = count
// 			mostrepeated = num
// 		}
// 	}
// 	return mostrepeated

// }

package main

import "fmt"

func main() {
	arr := []int{0, 0, 1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 5, 5}
	fmt.Println(mostRepeated(arr))
}

func mostRepeated(arr []int) []int {
	countMap := make(map[int]int)
	mostrepeated := []int{}
	maxcount := 0

	for _, num := range arr {
		countMap[num]++
	}
	for _, count := range countMap {
		if count > maxcount {
			maxcount = count
		}
	}

	for num, count := range countMap {
		if count == maxcount {
			mostrepeated = append(mostrepeated, num)
		}
	}
	return mostrepeated
}
