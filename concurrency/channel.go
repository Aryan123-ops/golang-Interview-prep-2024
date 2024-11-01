package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		sum := 0
		for i := 1; i <= 10; i++ {
			sum += i
		}
		c <- sum
	}()
	result := <-c
	fmt.Println(result)
}
