package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var (
		investmentAmount   float64
		years              float64
		expectedReturnRate = 5.5
	)

	// fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	// fmt.Print(("Years: "))
	outputText("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	// futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("futureValue:  %.2f\n", futureValue)
	formattedFRV := fmt.Sprintf("futureRealValue:  %.2f\n", futureRealValue)

	// fmt.Println("Future Value: ", futureValue)
	// fmt.Printf(`Future Value: %.2f\n
	// Future Real Value: %.2f\n`, futureValue, futureRealValue)
	// fmt.Println("Future Real Value: ", futureRealValue)

	fmt.Print(formattedFV + formattedFRV)
}

// ______________________________________________________________________
func outputText(text string) {
	fmt.Println(text)
}

// ______________________________________________________________________
func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	rfv = fv / math.Pow(1+inflationRate/100, years)

	return fv, rfv
	// return
}
