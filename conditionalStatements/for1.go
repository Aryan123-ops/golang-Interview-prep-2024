package main

import "fmt"

// func main() {
// 	num := 10
// 	for i := 0; i <= num; i++ {
// 		fmt.Printf("%d ", i)
// 	}
// }

// // for printng number if i becomes 5
// func main() {
// 	num := 10
// 	for i := 1; i <= num; i++ {
// 		if i == 5 {
// 			break
// 		}
// 		fmt.Printf("%d ", i)
// 	}
// 	fmt.Println("after break")
// }

// use of continue in for loop
// func main() {
// 	for i := 1; i <= 10; i++ {
// 		if i%2 == 0 {
// 			continue
// 		}
// 		fmt.Printf("%d ", i)
// 	}
// }

// addition of numbers from 1 to 10
func main() {
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
