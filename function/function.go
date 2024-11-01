package main

import "fmt"

func calculateBill(price, no int) int {
	totalPrice := price * no
	return totalPrice
}

func main() {
	billToBePay := calculateBill(20, 10)
	fmt.Println("Total Price", billToBePay)
}
