package main

import "fmt"

type Queue[T any] struct {
	elements []T
}

type QueueInterface[T any] interface {
	Enqueue(element T)
	Dequeue() (T, bool)
}

func (q *Queue[T]) Enqueue(element T) {
	q.elements = append(q.elements, element)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.elements) == 0 {
		var zero T
		return zero, false
	}
	elem := q.elements[0]
	q.elements = q.elements[1:]
	return elem, true
}

func main() {
	var queue Queue[string]
	queue.Enqueue("hi")
	queue.Enqueue("bye")
	fmt.Println(queue.Dequeue()) // Output: 10, true
	// for {
	//     elem, ok := queue.Dequeue()
	//     if !ok {
	//         break
	//     }
	//     fmt.Println(elem)
	// }
	// fmt.Println(queue.Dequeue()) // Output: 20, true
	// fmt.Println(queue.Dequeue()) // Output: 0, false
}
