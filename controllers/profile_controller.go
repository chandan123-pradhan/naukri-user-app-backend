package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"naurki_app_backend.com/services"
	"naurki_app_backend.com/utils"
)





func UpdateEmploymentHistory(w http.ResponseWriter, r *http.Request) {
	// Extract JWT Token from Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		respondWithJSON(w, http.StatusUnauthorized, "Authorization token is required", map[string]interface{}{})
		return
	}

	// Extract the Bearer token
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		respondWithJSON(w, http.StatusUnauthorized, "Invalid token format", map[string]interface{}{})
		return
	}

	// Validate the token
	userID, err := utils.VerifyJWT(tokenString)
	if err != nil {
		respondWithJSON(w, http.StatusUnauthorized, err.Error(), map[string]interface{}{})
		return
	}

	// Parse the incoming request to get the updated employment history
	var req struct {
		EmploymentHistory []struct {
			EmployerName string `json:"employer_name"`
			Designation  string `json:"designation"`
			StartDate    string `json:"start_date"`
			EndDate      string `json:"end_date"`
		} `json:"employment_history"`
	}

	// Decode the incoming JSON request into the req object
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, "Invalid request format", map[string]interface{}{})
		return
	}

	// Validate employment history input
	if len(req.EmploymentHistory) == 0 {
		respondWithJSON(w, http.StatusBadRequest, "Employment history is required", map[string]interface{}{})
		return
	}

	// Call service to update the employment history for the user
	err = services.UpdateEmploymentHistory(userID, req.EmploymentHistory)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, "Failed to update employment history", map[string]interface{}{})
		return
	}

	// Respond with success
	respondWithJSON(w, http.StatusOK, "Employment history updated successfully", map[string]interface{}{
		"user_id":           userID,
		"employment_history": req.EmploymentHistory,
	})
}






func GetUserDetails(w http.ResponseWriter, r *http.Request) {
	// Extract JWT Token from Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		respondWithJSON(w, http.StatusUnauthorized, "Authorization token is required", map[string]interface{}{})
		return
	}

	// Extract the Bearer token
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		respondWithJSON(w, http.StatusUnauthorized, "Invalid token format", map[string]interface{}{})
		return
	}

	// Validate the token and extract the user ID
	userID, err := utils.VerifyJWT(tokenString)
	if err != nil {
		respondWithJSON(w, http.StatusUnauthorized, err.Error(), map[string]interface{}{})
		return
	}

	// Fetch the user's details including employment history
	userDetails, err := services.GetUserDetails(userID)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, err.Error(), map[string]interface{}{})
		return
	}

	// Respond with the user's details
	respondWithJSON(w, http.StatusOK, "User details fetched successfully", map[string]interface{}{
		"user": userDetails,
	})
}