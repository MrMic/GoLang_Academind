package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser := user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(),
	}

	outputUserDetals(&appUser)
}

// ______________________________________________________________________
func outputUserDetals(u *user) {
	// fmt.Printf("First name: %s\nLast name: %s\nBirthdate: %s\nCreated at: %s\n", u.firstName, u.lastName, u.birthdate, u.createdAt)
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.createdAt)
}

// ______________________________________________________________________
func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
