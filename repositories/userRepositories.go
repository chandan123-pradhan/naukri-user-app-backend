package repositories

import (
	"database/sql"
	"fmt"

	"naurki_app_backend.com/config"
	"naurki_app_backend.com/models"
)

// CreateUser inserts a new user into the database
func CreateUser(user *models.User) error {
	// Prepare SQL statement (using ? placeholders for MySQL)
	stmt := `INSERT INTO users (full_name, email_id, highest_qualification, password, mobile_number, skills, description,prefferedLocation,prefferedSalary, prefferedShift, employmentType , profile_image_url) 
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	// Execute the query
	result, err := config.DB.Exec(stmt, user.FullName, user.EmailID, user.HighestQualification, user.Password, user.MobileNumber,
		user.Skills, user.Description, user.PrefferedLocation, user.PrefferedSallary, user.PrefferedShift, user.EmploymentType, user.ProfileImageURL)
	if err != nil {
		return fmt.Errorf("could not insert user: %v", err)
	}

	// Get the last inserted ID (auto_increment value)
	userID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("could not retrieve last inserted ID: %v", err)
	}

	// Set the user ID to the newly inserted ID
	user.ID = int(userID)

	return nil
}

func GetUserByEmail(emailID string) (*models.User, error) {
	// Prepare the SQL query to fetch user details by email ID
	stmt := `SELECT id, full_name, email_id, highest_qualification, password, mobile_number,skills, description,prefferedLocation,prefferedSalary, prefferedShift, employmentType , profile_image_url 
			 FROM users WHERE email_id = ?`

	// Execute the query and store the result
	var user models.User
	err := config.DB.QueryRow(stmt, emailID).Scan(&user.ID, &user.FullName, &user.EmailID, &user.HighestQualification, &user.Password, &user.MobileNumber,
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

func UpdateEmployment(userID int, employerName, designation, startDate, endDate string) error {
	// Prepare SQL query to update or insert employment history
	stmt := `INSERT INTO user_employment_history (user_id, employer_name, designation, start_date, end_date) 
			 VALUES (?, ?, ?, ?, ?)`

	// Execute the query to insert new employment record
	_, err := config.DB.Exec(stmt, userID, employerName, designation, startDate, endDate)
	if err != nil {
		return fmt.Errorf("could not insert employment record: %v", err)
	}

	return nil
}
