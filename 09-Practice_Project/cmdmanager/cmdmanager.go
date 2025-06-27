// Package cmdmanager
package cmdmanager

import "fmt"

type CMDManager struct{}

// ReadLines - * INFO: METHOD -------------------------------------------------
func (cmd CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Please enter your prices. Confirm every price with Enter.")

	var prices []string
	for {
		var price string
		fmt.Print("Price: ")
		_, err := fmt.Scan(&price)
		if err != nil {
			if price == "" {
				break // Exit if no input is provided
			}
			fmt.Println("Error reading input, please try again.")
			continue // Prompt for input again
		}

		if price == "0" {
			break
		}

		prices = append(prices, price)
	}

	return prices, nil
}

// WriteResult - * INFO: METHOD --------------------------------------------
func (cmd CMDManager) WriteResult(data any) error {
	fmt.Println("Writing result:", data)
	return nil
}

// New - * INFO: Creates a new CMDManager instance.
func New() CMDManager {
	return CMDManager{}
}
