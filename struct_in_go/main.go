package main

import "fmt"

type contactInfo struct {
	email 	string
	zipCode int
}

type person struct {
	firstName	string
	lastName	string
	// Embedding structs
	//contact 	contactInfo
	// contactInfo	contactInfo
	contactInfo
}

func main() {
	// alex := person{"Alex", "Anderson"}
	// alex := person{firstName: "Alex", lastName: "Anderson"}

	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Printf("%+v\n", alex)

	// jim := person {
	// 	firstName: "Jim",
	// 	lastName: "Party",
	// 	contact: contactInfo{
	// 		email: "jim@gmail.com",
	// 		zipCode: 94000,
	// 	},
	// }
	jim := person {
		firstName: "Jim",
		lastName: "Party",
		contactInfo: contactInfo{
			email: "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// jimPointer := &jim
	// jimPointer.updateName("Jimmy")

	// jim.updateName("jimmy")
	// jim.print()

	// Pointer shortcut, no need jimPointer
	jim.updateName("jimmy")
	jim.print()

	name := "bill"
	namePointer := &name
	
	fmt.Println(&namePointer)
	printPointer(namePointer)
}

func printPointer(namePointer *string) {
	fmt.Println(&namePointer)
}

// func (p person) updateName(newFirstName string) {
// 	p.firstName = newFirstName
// }
// func (pointerToPerson *person) updateName(newFirstName string) {
// 	(*pointerToPerson).firstName = newFirstName
// }

func (p *person) updateName(newFirstName string) {
	// *p.firstname = newFirstName => invalid operation
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}