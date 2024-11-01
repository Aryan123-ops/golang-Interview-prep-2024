package main

import (
	"fmt"
	"sync"
)

func HelloFrom1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello fromm 1")
}

func HelloFrom2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello fromm 2")
}

func execute() {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go HelloFrom1(wg)
	go HelloFrom2(wg)

	wg.Wait()
}
func main() {
	execute()
	fmt.Println("Main exited!")
}

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	execute()
// 	fmt.Println("main executed")
// }

// func hello(done chan bool) {
// 	fmt.Println("Hello")
// 	done <- true
// }

// func bye() {
// 	fmt.Println("Bye")
// }

// func execute() {
// 	done := make(chan bool)
// 	go hello(done)
// 	<-done // Wait for hello to finish
// 	bye()
// }
