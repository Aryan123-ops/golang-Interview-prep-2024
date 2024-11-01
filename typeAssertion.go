// type assertion is the technique to chek if the interface contians the specific data type and retrieves the value from that interface.

package main

import "fmt"

var Data interface{} = 42

func main() {
	if value, ok := Data.(int); ok {
		fmt.Println("Data is an int", value)
		return
	} else if value, ok := Data.(string); ok {
		fmt.Println("Value is string", value)
		return
	} else if value, ok := Data.(float64); ok {
		fmt.Println("Value is float", value)
		return
	} else {
		fmt.Println("Not of any type")
	}
}
