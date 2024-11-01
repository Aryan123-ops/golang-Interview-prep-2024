package main

import "fmt"

func main() {
	arr := []int{1, 1, 2, 2, 3, 4, 5, 5, 5}

	countMap := make(map[int]int)

	for _, num := range arr {
		countMap[num]++
	}

	for k, num := range countMap {
		fmt.Printf("%d:%d\n", k, num)
	}
}
