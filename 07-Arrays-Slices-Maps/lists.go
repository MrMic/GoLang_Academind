package main

import "fmt"

func main() {
	productNames := [4]string{"A Book"}
	prices := [4]float64{10.99, 9.99, 45.99, 20.00}

	productNames[2] = "A Carpet"

	fmt.Println(prices)
	fmt.Println(productNames)

	fmt.Println("🪚 prices:", prices[2])

	// * INFO: SLICE => featuredPrices := prices[1:3]
	featuredPrices := prices[1:]
	highlightedPrices := featuredPrices[:1]
	fmt.Println("🪚 featuredPrices:", featuredPrices)
	fmt.Println("🪚 highlightedPrices:", highlightedPrices)

	featuredPrices[0] = 199.99
	fmt.Println("🪚 prices:", prices)

	fmt.Println("len: ", len(featuredPrices), "capacity: ", cap(featuredPrices))
	fmt.Println("len: ", len(highlightedPrices), "capacity: ", cap(highlightedPrices))

	highlightedPrices = highlightedPrices[:3]
	fmt.Println("🪚 highlightedPrices:", highlightedPrices)
	fmt.Println("len: ", len(highlightedPrices), "capacity: ", cap(highlightedPrices))
}
