package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	// Implement password hashing logic with bcrypt
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err

}
