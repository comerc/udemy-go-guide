package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// alex := person{"Alex", "Anderson"}
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	// jimPointer := &jim
	jim.updateName("jimmy")
	jim.print()
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p *person) print() {
	fmt.Printf("%+v", p)
}
