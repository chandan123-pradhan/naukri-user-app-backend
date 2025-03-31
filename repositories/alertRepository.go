package repositories

import (
	"fmt"

	"naurki_app_backend.com/config"
	"naurki_app_backend.com/models"
)

// GetEmploymentHistoryByUserID fetches the employment history for a given user ID
func CreateAlert(alertData models.CreateAlertModel, userId int, profilePic string) error {
	// Prepare SQL query to fetch the user's employment history
	stmt := `INSERT INTO user_alerts (jobTitle, skills, email, username, mobile, userId, location,description, profile_image_url) 
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	// Execute the query and scan the result into a slice of employment history
	result, err := config.DB.Exec(stmt, alertData.JobTitle, alertData.Skills,alertData.Email,alertData.UserName,alertData.MobileNo, userId,alertData.Location, alertData.Description, profilePic)
	if err != nil {
		return fmt.Errorf("could not insert user: %v", err)
	}

	// Get the last inserted ID (auto_increment value)
	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("could not retrieve last inserted ID: %v", err)
	}
	fmt.Println("new alert created ", id)
	return  nil
}
