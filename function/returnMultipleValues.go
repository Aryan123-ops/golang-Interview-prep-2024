// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	num, message := myfunc()
	fmt.Println(num)
	fmt.Println(message)
}

func myfunc() (int, string) {
	return 42, "Hello World"

}
