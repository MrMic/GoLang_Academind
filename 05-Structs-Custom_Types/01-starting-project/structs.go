package main

import (
	"errors"
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

// ______________________________________________________________________
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

// ______________________________________________________________________
func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// * INFO: ══ CONSTRUCTOR ═════════════════════════════════════════════════════
func newUser(firstName, lastName, birthdate string) (*user, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("first name, last name, and birthdate are required")
	}

	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

// * INFO: ══ MAIN ════════════════════════════════════════════════════════════
func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, err := newUser(userFirstName, userLastName, userBirthdate)
	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()
}

// ______________________________________________________________________
func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
