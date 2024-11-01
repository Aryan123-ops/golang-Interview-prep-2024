// polymorphism = "ability of the fucntion to appear more than one form or to be used in diffrent parts of the code for same purpose"

package main

import "fmt"

// Define struct that will implement shape interface
type Circle struct {
	radius float64
}

// Define struct that will implement shape interface
type Rectangle struct {
	width  float64
	height float64
}

// Define an interface that will have method signatures that will be used.
type Shape interface {
	Area() float64
}

// Implement the Area method for Circle
func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

// Implement the Area method for Rectangle
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func main() {
	// Create instances of Circle and Rectangle
	circle := Circle{radius: 1.5}
	rectangle := Rectangle{width: 2.5, height: 2.5}

	// Create a slice of Shape interface
	shapes := []Shape{circle, rectangle}

	// loop through slice to calculate diffrent shapes of different Struct
	for _, shape := range shapes {
		fmt.Printf("Area %f\n", shape.Area())
	}
}
