// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func First() {
	fmt.Println("first")
}

func Second() {
	fmt.Println("second")
}

func Third() {
	fmt.Println("third")
}

func main() {

	defer Third()
	defer First()
	defer Second()
}
