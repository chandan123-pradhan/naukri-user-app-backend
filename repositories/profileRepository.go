package repositories

import (
	"database/sql"
	"fmt"

	"naurki_app_backend.com/config"
	"naurki_app_backend.com/models"
)

// GetUserByID fetches a user's details by their ID
func GetUserByID(userID int) (*models.User, error) {
	// Prepare SQL query to fetch the user details
	stmt := `SELECT id, full_name, email_id, highest_qualification, mobile_number, 
	
	skills, description,prefferedLocation,prefferedSalary, prefferedShift, employmentType ,
	profile_image_url 
			 FROM users WHERE id = ?`

	// Execute the query and scan the result into the user model
	var user models.User
	err := config.DB.QueryRow(stmt, userID).Scan(&user.ID, &user.FullName, &user.EmailID, &user.HighestQualification, &user.MobileNumber,
		
		&user.Skills,&user.Description, &user.PrefferedLocation, &user.PrefferedSallary, &user.PrefferedShift,&user.EmploymentType,
		
		&user.ProfileImageURL)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to fetch user: %v", err)
	}

	return &user, nil
}

// GetEmploymentHistoryByUserID fetches the employment history for a given user ID
func GetEmploymentHistoryByUserID(userID int) ([]models.EmploymentDetail, error) {
	// Prepare SQL query to fetch the user's employment history
	stmt := `SELECT employer_name, designation, start_date, end_date
			 FROM user_employment_history WHERE user_id = ?`

	// Execute the query and scan the result into a slice of employment history
	rows, err := config.DB.Query(stmt, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch employment history: %v", err)
	}
	defer rows.Close()

	var employmentHistory []models.EmploymentDetail
	for rows.Next() {
		var employment models.EmploymentDetail
		err := rows.Scan(&employment.EmployerName, &employment.Designation, &employment.StartDate, &employment.EndDate)
		if err != nil {
			return nil, fmt.Errorf("failed to scan employment record: %v", err)
		}
		employmentHistory = append(employmentHistory, employment)
	}

	// Check for errors from iterating over the rows
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate over employment records: %v", err)
	}

	return employmentHistory, nil
}
