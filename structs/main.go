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
	fmt.Printf("Person: %+v\n", person1)

	var person2 person
	person2.firstName = "Jaime"
	person2.lastName = "Lannister"

	fmt.Printf("Person: %+v\n", person2)

	person3 := person{
		firstName: "Cersei",
		lastName:  "Lannister",
		contact: contactInfo{
			email:   "cersei.lannister@gmail.com",
			zipCode: 56659,
		},
	}
	fmt.Printf("Person: %+v\n", person3)
}
