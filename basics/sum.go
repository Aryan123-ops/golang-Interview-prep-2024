package main

import "fmt"

// func main() {
// 	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	result := Sum(array)
// 	fmt.Println(result)
// }

// func Sum(arr []int) int {
// 	sum := 0
// 	for i := 1; i <= len(arr); i++ {
// 		sum += i
// 	}
// 	return sum
// }

func main() {
	array := []int{1, 2, 3, 4, 5}
	sum := 0
	for i := 0; i <= len(array); i++ {
		sum += i
	}
	fmt.Println(sum)
}

// func main() {
// 	num := 10
// 	sum := 0
// 	for i := 0; i <= num; i++ {
// 		sum += i
// 	}
// 	fmt.Println(sum)
// }
