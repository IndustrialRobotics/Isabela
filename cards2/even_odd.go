package main

import "fmt"

func EvenOrOdd() {

	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, value := range nums {
		if value%2 == 0 {
			fmt.Println(i, "Its Even")
		} else {
			fmt.Println(i, "Its Odd")
		}
	}

}
