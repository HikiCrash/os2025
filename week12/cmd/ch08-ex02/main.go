package main

import (
	"fmt"

	"github.com/headfirstgo/magazine"
)

func main() {
	var s1 magazine.Subscriber
	var e1 magazine.Employee
	s1.Name = "Choi"
	e1.Name = "Lee"
	e1.Salary = 5000000
	e1.Address.City = "Incheon"
	s1.Address.City = "Seoul"
	fmt.Println(s1.Name)
	fmt.Println(e1.Name, e1.Salary)
	fmt.Println(e1.Address.City, s1.Address.City)
}
