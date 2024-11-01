package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 45}
	fmt.Println(SecondLargest(arr))
}

func SecondLargest(arr []int) int {
	large1 := 0
	large2 := 0

	large1 = arr[0]
	for i := 1; i <= len(arr)-1; i++ {
		if large1 < arr[i] {
			large2 = large1
			large1 = arr[i]
		} else if large2 < arr[i] {
			large2 = arr[i]
		}
	}
	return large2
}

// package main

// import "fmt"

// func main() {
// 	fmt.Println("Try programiz.pro")
// 	arr := []int{-2, -3, -4}
// 	fmt.Println(secondLargest(arr))
// }

// func secondLargest(arr []int) int {
// 	if len(arr) < 2 {
// 		return -1 // Handle case where there are less than 2 elements
// 	}

// 	large1 := arr[0]
// 	large2 := arr[0]

// 	for i := 1; i < len(arr); i++ {
// 		if arr[i] > large1 {
// 			large2 = large1
// 			large1 = arr[i]
// 		} else if arr[i] > large2 && arr[i] != large1 {
// 			large2 = arr[i]
// 		}
// 	}

// 	// If large2 is still the same as large1, it means all elements were the same
// 	if large1 == large2 {
// 		return -1 // No distinct second largest element
// 	}

// 	return large2
// }
