// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	var wg sync.WaitGroup
// 	ch := make(chan int, 1)
// 	wg.Add(2)
// 	go odd(ch, &wg)
// 	go even(ch, &wg)
// 	ch <- 1
// 	wg.Wait()
// }

// func even(ch chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 2; i <= 10; i += 2 {
// 		<-ch
// 		fmt.Println("Even number", i)
// 		ch <- 1
// 	}
// }

// func odd(ch chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 1; i <= 10; i += 2 {
// 		<-ch
// 		fmt.Println("Odd number", i)
// 		ch <- 1
// 	}
// }

package main

import "fmt"

func main() {
	countChan := make(chan bool)
	done := make(chan bool)
	num := 10
	go even(num, countChan, done)
	countChan <- true
	go odd(num, countChan, done)

	<-done
	close(countChan)
}

func even(num int, countChan, done chan bool) {
	for i := 1; i <= num; i++ {
		<-countChan
		if i%2 == 0 {
			fmt.Println("even ", i)
		}
		countChan <- true
	}
	done <- true
}

func odd(num int, countChan, done chan bool) {
	for i := 1; i <= num; i++ {
		<-countChan
		if i%2 != 0 {
			fmt.Println("odd", i)
		}
		countChan <- true
	}
	done <- true
}

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	oddCh := make(chan bool)
// 	evenCh := make(chan bool)
// 	var wg sync.WaitGroup

// 	wg.Add(2)
// 	go printOdd(oddCh, evenCh, &wg)
// 	go printEven(evenCh, oddCh, &wg)

// 	go func() {
// 		wg.Wait()
// 		close(oddCh)
// 		close(evenCh)
// 	}()

// 	oddCh <- true
// 	wg.Wait() // Wait for both goroutines to finish
// }

// func printOdd(oddCh, evenCh chan bool, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 1; i <= 9; i += 2 {
// 		if _, ok := <-oddCh; !ok {
// 			return
// 		}
// 		fmt.Println("ODD NO", i)
// 		evenCh <- true
// 	}
// }

// func printEven(evenCh, oddCh chan bool, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 2; i <= 10; i += 2 {
// 		if _, ok := <-evenCh; !ok {
// 			return
// 		}
// 		fmt.Println("EVEN NO", i)
// 		oddCh <- true
// 	}
// }
