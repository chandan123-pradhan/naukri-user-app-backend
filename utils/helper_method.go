package utils

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"time"
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


func CheckOTPExpiration(createdAt time.Time) (string, error) {
	
	currentTime := time.Now().UTC()
fmt.Println(currentTime)
fmt.Println(createdAt)
	// Calculate the absolute difference in seconds
	duration := math.Abs(currentTime.Sub(createdAt).Seconds())
	fmt.Println("Time difference in seconds:", duration)

	// Check if more than 60 seconds have passed (1 minute)
	if duration > 60 {
			return "", fmt.Errorf("OTP expired. Please request a new OTP")
	}

	return "OTP Valid", nil // Or whatever success return you use
}


// Helper function to respond with JSON in the desired structure
func RespondWithJSON(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// Set the response structure
	response := map[string]interface{}{
		"status":  "failure", // Default status is failure
		"message": message,
		"data":    data, // Send the provided data (or empty if nil)
	}

	// If there's no error, set the status to "success"
	if statusCode == http.StatusOK || statusCode == http.StatusCreated {
		response["status"] = "success"
	}

	// Encode the response as JSON
	json.NewEncoder(w).Encode(response)
}

