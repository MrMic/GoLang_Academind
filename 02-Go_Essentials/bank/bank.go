package main

import "fmt"

func main() {
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check your balance")
	fmt.Println("2. Make a deposit")
	fmt.Println("3. Make a withdrawal")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Enter your choice: ")

	fmt.Scan(&choice)

	fmt.Println("Your choise: ", choice)
}
