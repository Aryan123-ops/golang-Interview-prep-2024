// package main

// import (
// 	"fmt"
// )

// func main() {
// 	// Create a channel to communicate prime numbers
// 	primeChan := make(chan int)

// 	// Launch a goroutine to find prime numbers
// 	go findPrimes(10, primeChan)

// 	// Read from the channel and print the prime numbers
// 	for prime := range primeChan {
// 		fmt.Println(prime)
// 	}
// }

// func findPrimes(limit int, primeChan chan int) {
// 	defer close(primeChan) // Ensure the channel is closed once done

// 	for i := 1; i <= limit; i++ {
// 		if isPrime(i) {
// 			primeChan <- i // Send prime number to the channel
// 		}
// 	}
// 	close(primeChan)
// }

// func isPrime(n int) bool {
// 	if n <= 1 {
// 		return false
// 	}
// 	if n == 2 {
// 		return true
// 	}
// 	if n%2 == 0 {
// 		return false
// 	}
// 	for i := 3; i*i <= n; i += 2 {
// 		if n%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

package main

import "fmt"

func main() {

	primechan := make(chan int)

	go findprime(10, primechan)

	for prime := range primechan {
		fmt.Println(prime)
	}
}

func isprime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func findprime(num int, primechan chan int) {
	for i := 2; i <= num; i++ {
		if isprime(i) {
			fmt.Println(i)
		}
	}
	close(primechan)
}
