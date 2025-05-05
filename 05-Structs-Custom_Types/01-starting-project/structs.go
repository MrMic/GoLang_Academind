package main

import (
	"fmt"
	"time"
)

// ______________________________________________________________________
type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

// outputUserDetails outputs the details of a user to the console.
//
// This method prints the user's first name, last name, birthdate, and the account creation date
// in a human-readable format.
//
// Parameters:
//
//	None
//
// Returns:
//
//	None
//
// Example:
//
//	user := user{
//	    firstName: "John",
//	    lastName: "Doe",
//	    birthdate: "1990-01-01",
//	    createdAt: "2023-01-01",
//	}
//	user.outputUserDetails() // Prints: John Doe 1990-01-01 2023-01-01
func (u *user) outputUserDetails() {
	// fmt.Printf("First name: %s\nLast name: %s\nBirthdate: %s\nCreated at: %s\n", u.firstName, u.lastName, u.birthdate, u.createdAt)
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.createdAt)
}

func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func newUser(firstName, lastName, birthdate string) *user {
	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}
}

// * INFO: ══ MAIN ════════════════════════════════════════════════════════════
func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser := newUser(userFirstName, userLastName, userBirthdate)

	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()
}

// ______________________________________________________________________
func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
