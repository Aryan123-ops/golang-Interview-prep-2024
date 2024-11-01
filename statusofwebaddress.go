package main

import (
	"fmt"
	"net/http"
)

// func main() {
// 	addres := "https://go.dev/play/"

// 	response, _ := http.Get(addres)
// 	fmt.Println(response.StatusCode)
// }

func main() {
	urls := []string{"https://go.dev/play/", "https://www.facebook.com/"}

	for _, status := range urls {
		response, _ := http.Get(status)
		fmt.Println(response.Status)
	}
}
