package fakeData

import (
	"math"
	"math/rand"
	"time"
)

func GenerateData(max int, volatility float64, responseTime int) func() float64 {

	value := rand.Float64()
	return func() float64 {
		time.Sleep(time.Duration(responseTime) * time.Millisecond)
		rnd := 2 * (rand.Float64() - 0.7)
		change := volatility * rnd
		value += change
		return math.Max(0, value*float64(max))
	}
}
