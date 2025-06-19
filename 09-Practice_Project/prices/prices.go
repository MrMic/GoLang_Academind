package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

// LoadData - * INFO: METHOD ----------------------------------------------------
func (job *TaxIncludedPriceJob) LoadData() {
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

	prices := make([]float64, len(lines))
	for lineIndex, line := range lines {
		floatPrice, err := strconv.ParseFloat(line, 64)
		if err != nil {
			file.Close() // Close the file before returning
			fmt.Println("Error parsing float from line:", line, "Error:", err)
		}
		prices[lineIndex] = floatPrice
	}
	job.InputPrices = prices

	defer file.Close()
}

// Process - * INFO: Processes the input prices by applying the tax rate and storing the results.
func (job *TaxIncludedPriceJob) Process() {
	job.LoadData() // Load data from file

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		TaxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", TaxIncludedPrice)
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
