package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 1)
	var wg sync.WaitGroup
	wg.Add(1)
	go multipleof5(ch, &wg)

	ch <- 1

	wg.Wait()
}

func multipleof5(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		<-ch
		fmt.Println(i * 5)
		ch <- 1
	}
}
