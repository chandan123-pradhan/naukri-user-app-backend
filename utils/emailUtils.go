package utils

import (
	"regexp"
)

// IsValidEmail checks if the given email address is in a valid format
func IsValidEmail(email string) bool {
	// Simple regex for basic email validation
	re := regexp.MustCompile(`^[a-z0-9]+@[a-z0-9]+\.[a-z]+$`)
	return re.MatchString(email)
}
