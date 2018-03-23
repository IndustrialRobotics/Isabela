package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email string
	age   int
}

func main() {

	jim := person{
		firstName: "John",
		lastName:  "Prince",

		contactInfo: contactInfo{
			email: "john@prince.com",
			age:   458},
	}

	jim.print()

	jim.updateName("Angelica")
	jim.print()




}

func (p person) print() {
	fmt.Println(p)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
