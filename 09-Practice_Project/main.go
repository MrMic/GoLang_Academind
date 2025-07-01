package main

import (
	"fmt"

	"michaelchlon.fr/price-calculator/filemanager"
	"michaelchlon.fr/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		err := priceJob.Process()
		if err != nil {
			fmt.Printf("Error processing job with tax rate %.2f: %v\n", taxRate, err)
		}
	}
}
