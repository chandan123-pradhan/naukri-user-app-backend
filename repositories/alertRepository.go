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


func GetUserAlertsByUserID(userID int) ([]models.UserAlert, error) {
	query := `SELECT id, jobTitle, skills, email, username, mobile, userId, location, description, profile_image_url 
			  FROM user_alerts 
			  WHERE userId = ?`

	rows, err := config.DB.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("could not fetch alerts for user ID %d: %v", userID, err)
	}
	defer rows.Close()

	var alerts []models.UserAlert
	for rows.Next() {
		var alert models.UserAlert
		err := rows.Scan(
			&alert.ID,
			&alert.JobTitle,
			&alert.Skills,
			&alert.Email,
			&alert.UserName,
			&alert.MobileNo,
			&alert.UserID,
			&alert.Location,
			&alert.Description,
			&alert.ProfileImageUrl,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning alert row: %v", err)
		}
		alerts = append(alerts, alert)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return alerts, nil
}
