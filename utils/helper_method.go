package utils

import (
	"fmt"

)

// Validate registration input
func ValidateRegistrationInput(req struct {
	FullName             string `json:"full_name"`
	EmailID              string `json:"email_id"`
	HighestQualification string `json:"highest_qualification"`
	Password             string `json:"password"`
	ConfirmPassword      string `json:"confirm_password"`
	MobileNumber         string `json:"mobile_number"`
	PrefferedLocation    string `json:"preffered_location"`
	PrefferedSallary     string `json:"preffered_sallary"`
	PrefferedShift       string `json:"preffered_shift"`
	EmploymentType       string `json:"employment_type"`
	Description          string `json:"description"`
	Skills               string `json:"skills"`
}) error {

	if req.FullName == "" {
		return fmt.Errorf("Full name is required")
	}
	if req.EmailID == "" {
		return fmt.Errorf("Email ID is required")
	}
	
	if req.HighestQualification == "" {
		return fmt.Errorf("Highest qualification is required")
	}
	if req.MobileNumber == "" {
		return fmt.Errorf("Mobile number is required")
	}
	if req.Password == "" {
		return fmt.Errorf("Password is required")
	}

	if req.Password != req.ConfirmPassword {
		return fmt.Errorf("Passwords do not match")
	}

	// You can also validate mobile number and other fields here

	return nil
}
