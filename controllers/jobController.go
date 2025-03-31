package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"naurki_app_backend.com/services"
	"naurki_app_backend.com/utils"
)



// GetJobPosts handles the request to fetch all job posts
func GetJobPosts(w http.ResponseWriter, r *http.Request) {
	_, err := utils.ValidateToken(w, r)
	if err != nil {
		return // Error response has already been sent by validateToken
	}
	jobPosts, err := services.GetJobPosts()
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, "Failed to fetch job posts", nil)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, "Job posts fetched successfully", map[string]interface{}{
		"jobs": jobPosts, // List of job posts
	})
}

// ApplyJob handles the job application process
func ApplyJob(w http.ResponseWriter, r *http.Request) {
	// Validate JWT and get the user ID
	userID, err := utils.ValidateToken(w, r)
	if err != nil {
		return // Error response has already been sent by validateToken
	}

	// Parse the request body
	var req struct {
		JobId string `json:"job_id"`
	}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, "Invalid request format", nil)
		return
	}

	if len(req.JobId) == 0 {
		utils.RespondWithJSON(w, http.StatusBadRequest, "Job ID Not found", nil)
		return
	}

	// Check if the job ID is valid
	if !services.IsJobIdCorrect(req.JobId) {
		utils.RespondWithJSON(w, http.StatusNotFound, "Job Not Found, please check", nil)
		return
	}

	// Apply for the job
	err = services.ApplyJob(userID, req.JobId)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusConflict, err.Error(),  map[string]interface{}{
		
		});
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, "Job applied successfully.",  map[string]interface{}{
		
	});
}

// GetJobDetails handles the request to fetch details for a specific job
func GetJobDetails(w http.ResponseWriter, r *http.Request) {
	// Validate JWT and get the user ID
	userId, err := utils.ValidateToken(w, r)
	if err != nil {
		return // Error response has already been sent by validateToken
	}

	// Parse the request body
	var req struct {
		JobId string `json:"job_id"`
	}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, "Invalid request format", nil)
		return
	}

	if len(req.JobId) == 0 {
		utils.RespondWithJSON(w, http.StatusBadRequest, "Job ID Not found", nil)
		return
	}

	// Check if the job ID is valid
	if !services.IsJobIdCorrect(req.JobId) {
		utils.RespondWithJSON(w, http.StatusNotFound, "Job Not Found, please check", nil)
		return
	}

	// Print user ID for debugging
	// You can implement further logic to fetch job details and return them in the response.
	jobDetails, err := services.JobDetails(req.JobId, strconv.Itoa(userId))
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	// Respond with the job details and applicants
	utils.RespondWithJSON(w, http.StatusOK, "Job details fetched successfully", jobDetails)

}



func GetAppliedJobs(w http.ResponseWriter, r *http.Request){
	userId, err := utils.ValidateToken(w, r)
	if err != nil {
		return // Error response has already been sent by validateToken
	}

	jobPosts, err := services.GetAppliedJobs(userId)
	fmt.Println(err)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, "Failed to fetch Applied Jobs", nil)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, "Applied jobs fetched successfully", map[string]interface{}{
		"jobs": jobPosts, // List of job posts
	})

}