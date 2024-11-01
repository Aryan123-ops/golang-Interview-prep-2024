package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile("/home/aryan/Downloads/Resume- Girish.docx")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("File opened successfully", file)
	}
}
