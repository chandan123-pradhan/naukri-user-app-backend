package services

import (
	"fmt"
	"naurki_app_backend.com/models"
	"naurki_app_backend.com/repositories"
)





// UpdateEmploymentHistory updates the user's employment history
func UpdateEmploymentHistory(userID int, employmentHistory []struct {
	EmployerName string `json:"employer_name"`
	Designation  string `json:"designation"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
}) error {
	// For each employment history entry, insert or update it
	for _, emp := range employmentHistory {
		err := repositories.UpdateEmployment(userID, emp.EmployerName, emp.Designation, emp.StartDate, emp.EndDate)
		if err != nil {
			return fmt.Errorf("failed to update employment history: %v", err)
		}
	}
	return nil
}




// GetUserDetails fetches the user's details along with employment history
func GetUserDetails(userID int) (*models.UserDetailsResponse, error) {
	// Fetch user details from the database (this would include basic user info like name, email, etc.)
	user, err := repositories.GetUserByID(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user details: %v", err)
	}

	// Fetch employment history for the user
	employments, err := repositories.GetEmploymentHistoryByUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch employment history: %v", err)
	}

	// Return the user details along with employment history
	return &models.UserDetailsResponse{
		ID:                 user.ID,
		FullName:           user.FullName,
		EmailID:            user.EmailID,
		HighestQualification: user.HighestQualification,
		MobileNumber:       user.MobileNumber,
		ProfileImageURL:    user.ProfileImageURL,
		EmploymentHistory:  employments,
		Description: user.Description,
		Skills: user.Skills,
		EmploymentType: user.EmploymentType,
		PrefferedLocation: user.PrefferedLocation,
		PrefferedSallary: user.PrefferedSallary,
		PrefferedShift: user.PrefferedShift,
	}, nil
}




