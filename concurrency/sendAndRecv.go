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

// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

// package main
// import (
//     "fmt"
//     "time"
// )
// func main() {
//   ch := make(chan int, 2)

//  go func(){
//      for i := 1; i <= 5; i++ {
//          ch <-i
//          fmt.Println("Sent", i)
//          time.Sleep(500 *time.Millisecond)
//      }
//      close(ch)
//  }()
//  for v := range ch {
//      fmt.Println("Received",v)
//  }
// }
