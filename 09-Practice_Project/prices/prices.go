package prices

import (
	"fmt"

	"michaelchlon.fr/price-calculator/conversion"
	"michaelchlon.fr/price-calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOMAnager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`            // Tax rate to apply to the input prices
	InputPrices       []float64           `json:"input_prices"`        // Input prices to which the tax rate will be applied
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"` // Map to store the input prices and their corresponding tax-included prices
}

// LoadData - * INFO: METHOD ----------------------------------------------------
func (job *TaxIncludedPriceJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		return err // Return the error if reading lines fails
	}

	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		return err // Return the error if conversion fails
	}

	job.InputPrices = prices
	return nil // Return nil if everything is successful
}

// Process - * INFO: Processes the input prices by applying the tax rate and storing the results.
func (job *TaxIncludedPriceJob) Process(doneChan chan bool) {
	err := job.LoadData() // Load data from file
	if err != nil {
		// return err // Return the error if loading data fails
	}

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		TaxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", TaxIncludedPrice)
	}

	job.TaxIncludedPrices = result // Store the results in the job's map

	job.IOManager.WriteResult(job)
	doneChan <- true // * NOTE: Signal that the processing is done
}

// NewTaxIncludedPriceJob - * INFO: Constructs a new TaxIncludedPriceJob with the given tax rate and input prices.
func NewTaxIncludedPriceJob(iom iomanager.IOMAnager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
