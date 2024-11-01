package main

import "fmt"

func main() {
	sumeven := 0
	sumodd := 0
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, num := range arr {
		if num%2 == 0 {
			sumeven += num
		}
		if num%2 != 0 {
			sumodd += num
		}
	}
	fmt.Println(sumeven)
	fmt.Println(sumodd)
}
