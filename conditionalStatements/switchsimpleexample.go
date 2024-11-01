package main

import "fmt"

func main() {
	finger := 4
	fmt.Printf("Finger is %d ", finger)

	switch finger {
	case 1:
		fmt.Println("indicates thumb")
	case 2:
		fmt.Println("indicates index")
	case 3:
		fmt.Println("indicates middle")
	case 4:
		fmt.Println("indicates ring")
	case 5:
		fmt.Println("indicates pinky")
	default:
		fmt.Println("indicates incorrect finger")
	}
}
