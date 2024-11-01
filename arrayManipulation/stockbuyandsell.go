package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	maxprofitTiime := maxProfit(prices)
	fmt.Println(maxprofitTiime)
}

func maxProfit(prices []int) int {
	min, res := 0, 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < prices[min] {
			min = i
		}
		if prices[i]-prices[min] > res {
			res = prices[i] - prices[min]
		}
	}
	return res
}
