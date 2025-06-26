package prices

import (
	"fmt"

	"michaelchlon.fr/price-calculator/conversion"
	"michaelchlon.fr/price-calculator/filemanager"
)

type TaxIncludedPriceJob struct {
	IOManager         filemanager.FileManager `json:"-"`
	TaxRate           float64                 `json:"tax_rate"`            // Tax rate to apply to the input prices
	InputPrices       []float64               `json:"input_prices"`        // Input prices to which the tax rate will be applied
	TaxIncludedPrices map[string]string       `json:"tax_included_prices"` // Map to store the input prices and their corresponding tax-included prices
}

// LoadData - * INFO: METHOD ----------------------------------------------------
func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		fmt.Println(err)
		return // Exit if there's an error reading the file
	}

	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		fmt.Println(err)
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

	job.IOManager.WriteResult(job)
}

// NewTaxIncludedPriceJob - * INFO: Constructs a new TaxIncludedPriceJob with the given tax rate and input prices.
func NewTaxIncludedPriceJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   fm,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
