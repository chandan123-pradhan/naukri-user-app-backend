package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"naurki_app_backend.com/models"
	"naurki_app_backend.com/services"
	"naurki_app_backend.com/utils"
)

var validate = validator.New() // Declare globally

func CreateAlerts(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.ValidateToken(w, r)
	if err != nil {
		return // Token validation already sends an error response
	}

	var req models.CreateAlertModel

	// Decode JSON request
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, "Invalid request format", map[string]interface{}{})
		return
	}

	// Validate request data
	err = validate.Struct(req)
	if err != nil {
		validationErrors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors[err.Field()] = fmt.Sprintf("Field '%s' is required", err.Field())
		}
		utils.RespondWithJSON(w, http.StatusBadRequest, "Validation failed", validationErrors)
		return
	}

	userDetails, err := services.GetUserDetails(userId)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, err.Error(), map[string]interface{}{})
		return
	}
	// 🔹 If validation passes, create alert
	err= services.CreateAlerts(req, userId, userDetails.ProfileImageURL)
	if(err!=nil){
		fmt.Println(err.Error())
		utils.RespondWithJSON(w, http.StatusInternalServerError, "Couldn't added alert, please try after sometime", map[string]interface{}{})
	return
	}

	
	utils.RespondWithJSON(w, http.StatusOK, "Alert Created Successfully", map[string]interface{}{})
	
}


func GetAlerts(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.ValidateToken(w, r)
	if err != nil {
		return // Token validation already sends an error response
	}

	alerts, err := services.GetAlerts(userId)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, "Failed to fetch alerts", map[string]interface{}{})
		return
	}

	// Ensure alerts is an empty slice, not nil
	if alerts == nil {
		alerts = []models.UserAlert{}
	}

	utils.RespondWithJSON(w, http.StatusOK, "User alerts fetched successfully", map[string]interface{}{
		"alerts": alerts,
	})
}
