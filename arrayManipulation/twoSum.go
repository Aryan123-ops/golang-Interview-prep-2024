// package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello, 世界")
// 	arr := []int{7, 1, 3, 8, 5, 3, 5, 9, 1}
// 	target := 17
// 	fmt.Println(TwoSum(arr, target))
// }

// func TwoSum(arr []int, target int) (int, int) {
// 	complements := make(map[int]int)

// 	for _, num := range arr {
// 		complement := target - num
// 		if val, found := complements[complement]; found {
// 			return val, num
// 		}
// 		complements[num] = num
// 	}
// 	return -1, -1
// }

package main

import (
	"fmt"
	"sort"
	// "sort"
)

func main() {

	// arr := []int{7, 1, 3, 8, 5, 3, 5, 9, 1}
	arr := []int{7, 9, 4, 3}

	target := 13
	fmt.Println(TwoSum(arr, target))
}

// Return the elments that sums up to target

func TwoSum(arr []int, target int) (int, int) {
	sort.Ints(arr)
	left := 0
	right := len(arr) - 1
	for left < right {
		sum := arr[left] + arr[right]
		if sum == target {
			return arr[left], arr[right]
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return -1, -1
}

// return the index of the element in array that sums up to target .

// func twoSum(nums []int, target int) []int {
// 	numMap := make(map[int]int) // Step 1: Create a map to store numbers and their indices.
// 	// Step 2: Add each number and its index to the map.
// 	for i, num := range nums {
// 		numMap[num] = i
// 	}
// 	// Step 3: Check for each number's complement in the map.
// 	for i, num := range nums {
// 		if j, ok := numMap[target-num]; ok && i != j {
// 			return []int{i, j} // Step 4: Return the indices if the complement exists.
// 		}
// 	}
// 	// Step 5: If no complement is found, return nil.
// 	return nil
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{1, 2, 3, 4, 5}
// 	target := 5
// 	i, j := twosum(arr, target)
// 	fmt.Println(i, j) // prints the indices
// }

// func twosum(arr []int, target int) (int, int) {
// 	indexMap := make(map[int]int) // value -> index

// 	for i, num := range arr {
// 		complement := target - num
// 		if idx, ok := indexMap[complement]; ok {
// 			return idx, i
// 		}
// 		indexMap[num] = i
// 	}
// 	return -1, -1
// }
