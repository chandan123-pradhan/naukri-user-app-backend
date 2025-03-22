package utils

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the password using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash compares the given password with the hashed password
func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}


// isValidPhoneNumber checks if the phone number is valid
func IsValidPhoneNumber(phone string) bool {
	// India mobile number validation pattern
	var validPhoneRegex = regexp.MustCompile(`^[6-9]\d{9}$`)
	return validPhoneRegex.MatchString(phone)
}