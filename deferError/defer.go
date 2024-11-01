package main

import (
	"fmt"
	"time"
)

func main() {
	test()
}

func timetaken(start time.Time) {
	fmt.Printf("total time taken %f seconds\n", time.Since(start).Seconds())
}

func test() {
	start := time.Now()
	defer timetaken(start)
	time.Sleep(2 * time.Second)
	fmt.Println("sleep complete")
}
