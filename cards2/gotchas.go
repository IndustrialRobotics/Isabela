package main

import "fmt"

type people struct {
	firstName string
	lastName  string
}

//Reference VS Value Type.png
func main() {

	jim := people{firstName: "Colobos", lastName: "Escorial"}
	mySlice := []string{"Hi", "There", "How", "Are", "You"}

	updateSlice(mySlice)
	fmt.Println(mySlice)

	jim.updatePeople("Altamira")
	nicePrint(jim.firstName)

}

func updateSlice(s []string) {
	s[0] = "Bye"
}

func (p *people) updatePeople(s string) {
	p.firstName = s
}

func nicePrint(s string) {
	fmt.Println(s)
}
