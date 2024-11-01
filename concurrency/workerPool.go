package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	for j := 1; j <= 5; j++ {
		jobs <- j
	}

	for a := 1; a <= 5; a++ {
		<-results
	}
	close(results)
}

func worker(id int, jobs, results chan int) {
	for j := range jobs {
		fmt.Printf("worker %d got job %d\n", id, j)
		time.Sleep(2 * time.Second)
		fmt.Printf("worker %d has finished the job %d\n", id, j)
		results <- j * 2
	}
}
