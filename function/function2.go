package main

import "fmt"

func main() {

	area, perimeter := rectProps(10.2, 12.2)
	fmt.Printf("Area is %f , Perimater is %f ", area, perimeter)
}

func rectProps(length, breadth float64) (float64, float64) {
	area := length * breadth
	perimeter := (length * breadth) / 2

	return area, perimeter
}
