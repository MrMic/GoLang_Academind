package main

import (
	"fmt"

	"michaelchlon.fr/price-calculator/filemanager"
	"michaelchlon.fr/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)

		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[index], errorChans[index])
		// if err != nil {
		// 	fmt.Printf("Error processing job with tax rate %.2f: %v\n", taxRate, err)
		// }
	}

	for index := range taxRates {
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Printf("Error processing job with tax rate %.2f: %v\n", taxRates[index], err)
			}
		case <-doneChans[index]:
			fmt.Println("DONE!")
		}
	}
}
