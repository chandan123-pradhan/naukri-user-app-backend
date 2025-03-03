package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"naurki_app_backend.com/services"
	"naurki_app_backend.com/utils"
)

// Helper function for validating the JWT token and extracting the user ID
func validateToken(w http.ResponseWriter, r *http.Request) (int, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		respondWithJSON(w, http.StatusUnauthorized, "Authorization token is required", nil)
		return 0, fmt.Errorf("authorization token is required")
	}

	// Extract the Bearer token
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		respondWithJSON(w, http.StatusUnauthorized, "Invalid token format", nil)
		return 0, fmt.Errorf("invalid token format")
	}

	// Validate the token and get the user ID
	userID, err := utils.VerifyJWT(tokenString)
	if err != nil {
		respondWithJSON(w, http.StatusUnauthorized, err.Error(), nil)
		return 0, err
	}

	return userID, nil
}

// GetJobPosts handles the request to fetch all job posts
func GetJobPosts(w http.ResponseWriter, r *http.Request) {
	jobPosts, err := services.GetJobPosts()
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, "Failed to fetch job posts", nil)
		return
	}

	respondWithJSON(w, http.StatusOK, "Job posts fetched successfully", map[string]interface{}{
		"jobs": jobPosts, // List of job posts
	})
}

// ApplyJob handles the job application process
func ApplyJob(w http.ResponseWriter, r *http.Request) {
	// Validate JWT and get the user ID
	userID, err := validateToken(w, r)
	if err != nil {
		return // Error response has already been sent by validateToken
	}

	// Parse the request body
	var req struct {
		JobId string `json:"job_id"`
	}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, "Invalid request format", nil)
		return
	}

	if len(req.JobId) == 0 {
		respondWithJSON(w, http.StatusBadRequest, "Job ID Not found", nil)
		return
	}

	// Check if the job ID is valid
	if !services.IsJobIdCorrect(req.JobId) {
		respondWithJSON(w, http.StatusNotFound, "Job Not Found, please check", nil)
		return
	}

	// Apply for the job
	err = services.ApplyJob(userID, req.JobId)
	if err != nil {
		respondWithJSON(w, http.StatusConflict, err.Error(),  map[string]interface{}{
		
		});
		return
	}

	respondWithJSON(w, http.StatusOK, "Job applied successfully.",  map[string]interface{}{
		
	});
}

// GetJobDetails handles the request to fetch details for a specific job
func GetJobDetails(w http.ResponseWriter, r *http.Request) {
	// Validate JWT and get the user ID
	userId, err := validateToken(w, r)
	if err != nil {
		return // Error response has already been sent by validateToken
	}

	// Parse the request body
	var req struct {
		JobId string `json:"job_id"`
	}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, "Invalid request format", nil)
		return
	}

	if len(req.JobId) == 0 {
		respondWithJSON(w, http.StatusBadRequest, "Job ID Not found", nil)
		return
	}

	// Check if the job ID is valid
	if !services.IsJobIdCorrect(req.JobId) {
		respondWithJSON(w, http.StatusNotFound, "Job Not Found, please check", nil)
		return
	}

	// Print user ID for debugging
	// You can implement further logic to fetch job details and return them in the response.
	jobDetails, err := services.JobDetails(req.JobId, strconv.Itoa(userId))
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	// Respond with the job details and applicants
	respondWithJSON(w, http.StatusOK, "Job details fetched successfully", jobDetails)

}



// func GetAppliedJobs(w http.ResponseWriter, r *http.Request){
// 	userId, err := validateToken(w, r)
// 	if err != nil {
// 		return // Error response has already been sent by validateToken
// 	}

// }