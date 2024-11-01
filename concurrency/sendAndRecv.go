package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go send(ch, wg)

	go recv(ch, wg)
	wg.Wait()
	fmt.Println("main exited!")
}

func send(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	msg := "hi"
	fmt.Println("sent message:", msg)

	ch <- msg
}

func recv(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	msg := <-ch
	fmt.Println("received message:", msg)

}
