package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	person1 := person{firstName: "Tyrion", lastName: "Lannister"}
	fmt.Printf("Person: %+v\n", person1)

	var person2 person
	person2.firstName = "Jaime"
	person2.lastName = "Lannister"

	fmt.Printf("Person: %+v\n", person2)

}
