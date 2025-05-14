package user

import (
	"errors"
	"fmt"
	"time"
)

// ______________________________________________________________________
type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

// ______________________________________________________________________
func (u *User) OutputUserDetails() {
	// fmt.Printf("First name: %s\nLast name: %s\nBirthdate: %s\nCreated at: %s\n", u.firstName, u.lastName, u.birthdate, u.createdAt)
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.createdAt)
}

// ______________________________________________________________________
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

// * INFO: ══ CONSTRUCTOR ADMIN ═════════════════════════════════════════════════════
func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthdate: "---",
			createdAt: time.Now(),
		},
	}
}

// * INFO: ══ CONSTRUCTOR ═════════════════════════════════════════════════════
func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("first name, last name, and birthdate are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}
