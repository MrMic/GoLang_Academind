package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ... do something awesome with that gathered data!

	outputUserDetals(firstName, lastName, birthdate)
}

// ______________________________________________________________________
func outputUserDetals(firstName, lastName, birthdate string) {
	fmt.Println(firstName, lastName, birthdate)
}

// ______________________________________________________________________
func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
