package services

import (
	"naurki_app_backend.com/models"
	"naurki_app_backend.com/repositories"
)



func CreateAlerts(createAlertModel models.CreateAlertModel, userId int, profilePic string)error{
	
	response := repositories.CreateAlert(createAlertModel, userId, profilePic)

	// You can do additional logic after calling ApplyJob if needed
	// For now, just return the message received from ApplyJob
	return response
	 
}


// GetAlerts retrieves alerts for a given user ID
func GetAlerts(userId int) ([]models.UserAlert, error) {
	alerts, err := repositories.GetUserAlertsByUserID(userId)
	if err != nil {
		return nil, err
	}
	return alerts, nil
}