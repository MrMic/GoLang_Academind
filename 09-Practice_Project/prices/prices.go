package prices

import (
	"bufio"
	"fmt"
	"os"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

// LoadData - * INFO: Method
func (job TaxIncludedPriceJob) LoadData() {
	file, err := os.Open("prices.txt")
	if err != nil { // Handle error
		fmt.Println("Error opening file:", err)
		return
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil { // Handle error
		// * INFO: Handle error
		fmt.Println("Error reading text file:", err)
		file.Close() // Close the file before returning
		return
	}

	defer file.Close()
}

func (job TaxIncludedPriceJob) Process() {
	result := make(map[string]float64)

	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}
	fmt.Println(result)
}

// NewTaxIncludedPriceJob - * INFO: Constructs a new TaxIncludedPriceJob with the given tax rate and input prices.
func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices:       []float64{10, 20, 30},
		TaxRate:           taxRate,
		TaxIncludedPrices: make(map[string]float64),
	}
}
