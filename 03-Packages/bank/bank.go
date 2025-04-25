package main

import (
	"fmt"

	"michaelchlon.com/bank/fileops"
)

var accountBalanceFile = "balance.txt"

// * INFO:  ════════════════════════════ MAIN ═════════════════════════
func main() {
	accountBalance, err := fileops.GetFloatFromFile(accountBalanceFile) // accountBalance := getBalanceFromFile()
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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
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
				fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			}

		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thaks for choosing Go Bank!")
			// break loopLabel
			return

		}
	}
}
