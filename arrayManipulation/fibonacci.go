package main

import "fmt"

func main() {
	fmt.Println(fibonacci(7))
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main
import "fmt"

func main() {
  fmt.Println("Try programiz.pro")
  fmt.Println(factorial(7))
}
func factorial(n int) int {
    if n == 0 {
        return 1
    }else {
        return factorial(n-1) + factorial(n-2)
    }
}
