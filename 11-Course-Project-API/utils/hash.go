package utils

import "golang.org/x/crypto/bcrypt"

// * NOTE: HashPassword hashes the password using bcrypt and returns the hashed password as a string
func HashPassword(password string) (string, error) {
	// Implement password hashing logic with bcrypt
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// * NOTE: CheckPasswordHash checks if the provided password matches the hashed password
func CheckPasswordHash(password, hashedPassword string) bool {
	// Implement password hash checking logic with bcrypt
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
