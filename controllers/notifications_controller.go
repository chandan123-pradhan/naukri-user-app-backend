package controllers

import (
	"encoding/json"

	"net/http"

	"naurki_app_backend.com/models"
	"naurki_app_backend.com/services"
	"naurki_app_backend.com/utils"
)

func SendNotification(w http.ResponseWriter, r *http.Request) {
	var req models.SendNotificationRequest

	// Parse the JSON request body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, "Invalid JSON body", nil)
		return
	}

	// Basic validation
	if req.FcmToken == "" || req.Title == "" || req.Body == "" {
		utils.RespondWithJSON(w, http.StatusBadRequest, "All fields (fcm_token, title, body) are required", nil)
		return
	}

	// Call the service to send the notification
	err := services.SendNotificationToToken(req.FcmToken, req.Title, req.Body)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, "Failed to send notification", nil)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, "Notification sent successfully", nil)
}
