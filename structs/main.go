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
	// 1. Turn an address into value then use *address
	// 2. Turn a value into address then use &value
	fmt.Printf("main() -> Person3 Origin Address: %p\n", &person3)
	/* Pointer Shortcut:
	   --> person3 is type of person, NOT type of *person as the receiver func expects
	   --> actually we need to type: (&person3).updateName("newName")
	   --> BUT this is working - it's a pointer shortcut
	*/
	person3.updateName("Snake")
	person3.print()
}

// !!! func (var *anyType) == is NOT dereferencing a variable, its a descriptive way to say this is type of pointer to a person
func (p *person) updateName(newFirstName string) {
	fmt.Printf("updateName() -> Address of the pointer: %p\n", &p)
	fmt.Printf("updateName() -> Address of the person itself: %p\n", &(*p))
	// !!! here (*p) is really dereferencing to get the actual value - NOT THE SAME AS IN THE SIGNATURE
	(*p).firstName = newFirstName // (*p) == dereferencing, so that I get the value from the address

}

func (p person) print() {
	fmt.Printf("Person: %+v\n", p)
}
