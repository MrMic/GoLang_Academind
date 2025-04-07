package main

import "math"

func main() {
	investmentAmount := 1000.00
	expectedReturnRate := 5.5
	years := 10.0

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	println(futureValue)
}
