package main

import (
	"Isabela/AndyPackage/calc"
	"fmt"
	"Isabela/AndyPackage/calc/convert"
)

func main() {

	doSum := calc.AddNum(45, 789)

	fmt.Println(doSum, calc.SubstractNum(5458, 56))
	fmt.Println(convert.SwapCase("hola a todos los fan de golang"))

}
