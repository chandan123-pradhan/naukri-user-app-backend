package controllers

import (
	"encoding/json"
	"net/http"

	"naurki_app_backend.com/services"
)


func SearchJobByTitle(w http.ResponseWriter, r *http.Request){
	_, err := validateToken(w, r)
	if err != nil {
		return // Error response has already been sent by validateToken
	}


	var req struct {
		JobTitle string `json:"title"`
	}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, "Invalid request format", nil)
		return
	}
	jobPosts, err := services.SearchJobByTitle(req.JobTitle)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, "Failed to fetch job posts", nil)
		return
	}

	respondWithJSON(w, http.StatusOK, "Job posts fetched successfully", map[string]interface{}{
		"jobs": jobPosts, // List of job posts
	})
}