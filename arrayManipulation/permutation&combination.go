package main

import (
	"fmt"
)

// Factorial function to compute factorial of a number
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

// Function to calculate permutation (nPn)
func permutation(n int) int {
	return factorial(n) // nPn = n!
}

// Function to calculate combination (nCn)
func combination(n int) int {
	return 1 // nCn = 1
}

func main() {
	n := 5

	// Calculate and print Permutation
	perm := permutation(n)
	fmt.Printf("Permutation (nPn) for n = %d is %d\n", n, perm)

	// Calculate and print Combination
	comb := combination(n)
	fmt.Printf("Combination (nCn) for n = %d is %d\n", n, comb)
}
