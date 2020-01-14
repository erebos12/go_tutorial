package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	person1 := person{firstName: "Tyrion", lastName: "Lannister"}
	person1.print()

	var person2 person
	person2.firstName = "Jaime"
	person2.lastName = "Lannister"
	person2.print()

	person3 := person{
		firstName: "Cersei",
		lastName:  "Lannister",
		contact: contactInfo{
			email:   "cersei.lannister@gmail.com",
			zipCode: 56659,
		},
	}
	person3.print()
}

func (p person) print() {
	fmt.Printf("Person: %+v\n", p)
}
