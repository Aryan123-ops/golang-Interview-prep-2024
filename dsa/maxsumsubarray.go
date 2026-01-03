package main

import "fmt"

func main() {
	arr := []int{2, 1, 5, 1, 3, 2}
	fmt.Println("Max sum of subarray of size 3 is:", maxsumarray(arr, 3))
}

func maxsumarray(nums []int, k int) int {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxsum := sum

	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]
		if sum > maxsum {
			maxsum = sum
		}
	}
	return maxsum
}
