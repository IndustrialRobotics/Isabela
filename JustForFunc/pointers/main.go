package main

import (
	"fmt"
)

type data struct {
	A string
	B string
	C uint
}

func (d *data) changeA(a string) {
	d.A = a
}

func (d data) doMaths() uint {
	return d.C * 100
}

func main() {
	myData := data{}

	myData.A = "Hello"
	myData.B = "World"
	myData.C = 42

	fmt.Printf("%+v\n", myData)

	myData.A = "Goodbye"
	fmt.Printf("%+v\n", myData)

	myData.changeA("Starbucks")
	myData.doMaths()
	fmt.Printf("%+v\n", myData)
	fmt.Println(myData.doMaths())
}
