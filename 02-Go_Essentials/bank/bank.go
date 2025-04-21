package main

import (
	"fmt"
	"os"
	"strconv"
)

var accountBalanceFile = "balance.txt"

// ______________________________________________________________________
func getBalanceFromFile() float64 {
	fmt.Println("Reading balance from file...")
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return 0
	}

	balance, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		fmt.Println("Error parsing balance:", err)
		return 0
	}

	return balance
}

// ______________________________________________________________________
func writeBalanceToFile(balance float64) {
	fmt.Println("Writing balance to file:", balance)
	balanceText := fmt.Sprintf("%.2f", balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0o644)
}

// * INFO:  ════════════════════════════ MAIN ═════════════════════════
func main() {
	accountBalance := getBalanceFromFile()

	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check your balance")
		fmt.Println("2. Make a deposit")
		fmt.Println("3. Make a withdrawal")
		fmt.Println("4. Exit")

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
			writeBalanceToFile(accountBalance)
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
				writeBalanceToFile(accountBalance)
			}

		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thaks for choosing Go Bank!")
			// break loopLabel
			return

		}
	}
}
