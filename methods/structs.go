package main

import "fmt"

func main() {

	user1 := users{
		Name:   "Aryan",
		Age:    18,
		Salary: 500,
	}

	user2 := users{Name: "Anuj", Age: 18, Salary: 500}
	fmt.Println(user1, user2)
}

type users struct {
	Name   string
	Age    int
	Salary int
}
