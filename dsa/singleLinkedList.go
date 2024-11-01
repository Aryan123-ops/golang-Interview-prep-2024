// // package main

// // import "fmt"

// // // Comparable is a constraint that requires the type to be comparable
// // type Comparable interface {
// // 	comparable
// // }

// // // Node defines a node in the linked list
// // type Node[T Comparable] struct {
// // 	Value T
// // 	Next  *Node[T]
// // }

// // // LinkedList defines the linked list structure
// // type LinkedList[T Comparable] struct {
// // 	Head *Node[T]
// // }

// // // Add appends a new node with the given value to the end of the list
// // func (list *LinkedList[T]) Add(value T) {
// // 	newNode := &Node[T]{Value: value}
// // 	if list.Head == nil {
// // 		list.Head = newNode
// // 		return
// // 	}
// // 	current := list.Head
// // 	for current.Next != nil {
// // 		current = current.Next
// // 	}
// // 	current.Next = newNode
// // }

// // // Print prints all elements of the linked list
// // func (list *LinkedList[T]) Print() {
// // 	current := list.Head
// // 	for current != nil {
// // 		fmt.Print(current.Value, " ")
// // 		current = current.Next
// // 	}
// // 	fmt.Println()
// // }

// // // Remove deletes the first node with the specified value
// // func (list *LinkedList[T]) Remove(value T) {
// // 	if list.Head == nil {
// // 		return
// // 	}
// // 	if list.Head.Value == value {
// // 		list.Head = list.Head.Next
// // 		return
// // 	}
// // 	current := list.Head
// // 	for current.Next != nil {
// // 		if current.Next.Value == value {
// // 			current.Next = current.Next.Next
// // 			return
// // 		}
// // 		current = current.Next
// // 	}
// // }

// // // Find returns true if a node with the specified value exists
// // func (list *LinkedList[T]) Find(value T) bool {
// // 	current := list.Head
// // 	for current != nil {
// // 		if current.Value == value {
// // 			return true
// // 		}
// // 		current = current.Next
// // 	}
// // 	return false
// // }

// // func main() {
// // 	list := &LinkedList[int]{}

// // 	list.Add(10)
// // 	list.Add(20)
// // 	list.Add(30)
// // 	list.Print() // Output: 10 20 30

// // 	list.Remove(20)
// // 	list.Print() // Output: 10 30

// // 	fmt.Println(list.Find(10)) // Output: true
// // 	fmt.Println(list.Find(20)) // Output: false
// // }

package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type List struct {
	head *Node
}

func (l *List) add(value int) {
	newNode := &Node{data: value}

	if l.head == nil {
		l.head = newNode
		return
	}

	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}

	curr.next = newNode
}

func (l *List) remove(value int) {
	if l.head == nil {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		return
	}

	curr := l.head
	for curr.next != nil && curr.next.data != value {
		curr = curr.next
	}

	if curr.next != nil {
		curr.next = curr.next.next
	}
}

func main() {
	list := &List{}
	list.add(1)
	list.add(2)
	list.add(3)
	list.add(4)

	fmt.Println("Initial List: ")
	printList(list)

	list.remove(2)
	fmt.Println("List after removing 2: ")
	printList(list)

	list.remove(4)
	fmt.Println("List after removing 4: ")
	printList(list)
}

func printList(l *List) {
	curr := l.head
	for curr != nil {
		fmt.Printf("%d ", curr.data)
		curr = curr.next
	}
	fmt.Println()
}
