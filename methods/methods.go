package main

import "fmt"

func main() {
	emp1 := Employee{
		Name:     "Aryan",
		Currency: "$",
		Salary:   12000,
	}
	emp1.DisplaySalry()
}

type Employee struct {
	Name     string
	Currency string
	Salary   int
}

func (e Employee) DisplaySalry() {
	fmt.Printf("salary of %s is %d%s", e.Name, e.Salary, e.Currency)
	return
}

// func main() {
// 	user := User{
// 		Name:    "Aryan",
// 		Address: "dsbnvdb",
// 	}
// 	user.UserDetail()
// }

// func (u User) UserDetail() {
// 	fmt.Printf("Employee Name is %s and address is %s", u.Name, u.Address)
// 	return
// }

// type User struct {
// 	Name    string
// 	Address string
// }
