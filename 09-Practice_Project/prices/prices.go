package prices

import (
	"fmt"

	"michaelchlon.fr/price-calculator/conversion"
	"michaelchlon.fr/price-calculator/filemanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string
}

// LoadData - * INFO: METHOD ----------------------------------------------------
func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := filemanager.ReadLines("prices.txt")
	if err != nil {
		fmt.Println("Error reading lines from file:", "Error:", err)
		return // Exit if there's an error reading the file
		// This will prevent further processing if the file cannot be read
	}

	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		fmt.Println("Error parsing float from line:", "Error:", err)
		return // Exit if there's an error parsing a line
	}

	job.InputPrices = prices
}

// Process - * INFO: Processes the input prices by applying the tax rate and storing the results.
func (job *TaxIncludedPriceJob) Process() {
	job.LoadData() // Load data from file

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		TaxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", TaxIncludedPrice)
	}

	job.TaxIncludedPrices = result // Store the results in the job's map

	filemanager.WriteJSON(fmt.Sprintf("result_%.0f.json", job.TaxRate*100), job.TaxIncludedPrices)
}

// NewTaxIncludedPriceJob - * INFO: Constructs a new TaxIncludedPriceJob with the given tax rate and input prices.
func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
