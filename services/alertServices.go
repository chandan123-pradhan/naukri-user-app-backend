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