package main

import "fmt"

func main() {
	arr := []int{1, 1, 1, 2, 2, 2, 3, 3}
	newArray := alloweElementOnlyTwice(arr)
	fmt.Println(newArray)
}

func alloweElementOnlyTwice(arr []int) []int {
	countmap := make(map[int]int)
	removeThrice := []int{}

	for _, num := range arr {
		if countmap[num] < 2 {
			removeThrice = append(removeThrice, num)
			countmap[num]++
		}
	}
	return removeThrice
}
