// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a interface{} = "42"

	b := a
	fmt.Println(b)
	// fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeFor[any]())
}
