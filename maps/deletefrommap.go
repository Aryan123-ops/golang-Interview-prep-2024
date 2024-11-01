// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main

import "fmt"

func main() {
	fmt.Println("Try programiz.pro")
	datamap := map[string]string{
		"Aryan": "Noida",
		"Anuj":  "Delhi",
		"Akash": "Sagarpur",
	}
	name := "Akash"
	fmt.Println(datamap, "Before deleting!")
	delete(datamap, name)
	fmt.Println("After Delete!", datamap)
}
