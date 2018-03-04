package main

import (
	"Isabela/newFakeDataFunc"
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())
	for {
		fk := fakeData.GenerateData(545, 0.2, 1000)
		if fk() > 0 {
			fmt.Println(fk())
		}
	}
}
