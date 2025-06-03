package main

import "fmt"

func main() {
	var stack Stack[int]
	stack.Push(10) // inserting element to stack at first time.
	stack.Push(20) // inserting element to stack second time.
	fmt.Println(stack.Pop())
	for { // for loop to print all elements from stack
		elem, ok := stack.Pop()
		if !ok {
			break
		}
		fmt.Println(elem)
	}
}

type Stack[T any] struct {
	elements []T
}

type StackInterface[T any] interface {
	Push(element T)
	Pop() (T, bool)
}

func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	elem := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return elem, true
}
