// package main

// import "fmt"

// func main() {
// 	x := []int{1, 2, 3, 4, 5} // len=5, cap=5
// 	fmt.Println(len(x))
// 	fmt.Println(cap(x))
// 	x = append(x, 6) //  len=6, cap=12
// 	fmt.Println(len(x))
// 	fmt.Println(cap(x))
// 	x = append(x, 7) //  len=7, cap=14
// 	fmt.Println(len(x))
// 	fmt.Println(cap(x))
// 	a := x[4:]
// 	fmt.Println(a) //a={}, len=0, cap=0
// 	fmt.Println(len(a))
// 	fmt.Println(cap(a))
// 	y := alterSlice(a) //  len=2, cap=2

// 	fmt.Println(x) // x= {1, 2, 3, 4, 5,6,7} len=7, cap=14
// 	fmt.Println(len(x))
// 	fmt.Println(cap(x))
// 	fmt.Println(y) // {10,11}
// 	fmt.Println(len(y)) // y= {10,11},len=2, cap=2
// 	fmt.Println(cap(y))

// }

package main

import "fmt"

func main(){
	a := []int{1,2,3,4,5}
	b := []int{0,0,0,0,0,0}
	a = append(a, b...)
	fmt.Println(len(a))
	fmt.Println(cap(a))

	fmt.Println("ghjkl")
	fmt.Println(len(b))
	fmt.Println(len(b))

}