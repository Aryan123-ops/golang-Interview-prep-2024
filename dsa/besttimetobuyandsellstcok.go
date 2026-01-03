package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println("Max profit is:", maxprofit(prices))
}

func maxprofit(prices []int) int {
	minprice := prices[0]
	profit := 0

	for _, p := range prices[1:] {
		if p < minprice {
			minprice = p
		}
		if p-minprice > profit {
			profit = p - minprice
		}
	}
	return profit
}
