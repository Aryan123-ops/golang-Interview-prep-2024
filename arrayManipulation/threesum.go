// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	arr := []int{1, 2, 3, 4}
// 	target := 7
// 	fmt.Println(ThreeSum(arr, target))
// }

// func ThreeSum(arr []int, target int) (int, int, int) {
// 	sort.Ints(arr)
// 	n := len(arr)

// 	for i := 0; i < n-2; i++ {
// 		left := i + 1
// 		right := n - 1
// 		for left < right {
// 			sum := arr[i] + arr[left] + arr[right]
// 			if sum == target {
// 				return arr[i], arr[left], arr[right]
// 			} else if sum < target {
// 				left++
// 			} else {
// 				right--
// 			}
// 		}
// 	}
// 	return -1, -1, -1
// }

package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{2, 4, 3, 0}
	target := 5
	fmt.Println(threesum(arr, target))
}

func threesum(arr []int, target int) (int, int, int) {
	sort.Ints(arr)
	n := len(arr)

	for i := 0; i < n-2; i++ {
		left := i + 1
		right := n - 1
		for left < right {
			sum := arr[i] + arr[left] + arr[right]
			if sum == target {
				return arr[i], arr[left], arr[right]
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}
	return -1, -1, -1
}
