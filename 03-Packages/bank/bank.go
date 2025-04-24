package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var accountBalanceFile = "balance.txt"

// ______________________________________________________________________
func getFloatFromFile(fileName string) (float64, error) {
	fmt.Println("Reading balance from file...")
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 1000, errors.New("failed to find file")
	}

	value, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return 0, errors.New("failed to parse stored value")
	}

	return value, nil
}

// ______________________________________________________________________
func writeFloatToFile(value float64, fileName string) {
	fmt.Println("Writing balance to file:", value)
	balanceText := fmt.Sprintf("%.2f", value)
	os.WriteFile(fileName, []byte(balanceText), 0o644)
}

// * INFO:  ════════════════════════════ MAIN ═════════════════════════
func main() {
	accountBalance, err := getFloatFromFile(accountBalanceFile) // accountBalance := getBalanceFromFile()
	if err != nil {
		fmt.Println("ERROR: ", err)
		fmt.Println("------------------")
		// panic("Can't continue")
	}

	fmt.Println("Welcome to Go Bank!")

	for {
		presentOptions()

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Enter the deposit amount: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid deposit amount!")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Your updated balance is: ", accountBalance)
			writeFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			var withdrawalAmount float64
			fmt.Print("Enter the withdrawal amount: ")
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 || withdrawalAmount > accountBalance {
				fmt.Println("withdrawal must be >= 0 or Insufficient funds!")
				continue
			} else {
				accountBalance -= withdrawalAmount
				fmt.Println("Your updated balance is: ", accountBalance)
				writeFloatToFile(accountBalance, accountBalanceFile)
			}

		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thaks for choosing Go Bank!")
			// break loopLabel
			return

		}
	}
}
