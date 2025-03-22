package services

import (
	"fmt"

	"naurki_app_backend.com/models"
	"naurki_app_backend.com/repositories"
)

// GetJobPosts fetches all job posts from the repository
func SearchJobByTitle(title string) ([]models.JobPost, error) {
	// Fetch the job posts from the repository
	jobPosts, err := repositories.GetJobByTitle(title)
	if err != nil {
		// Log the error and return a more detailed error message
		fmt.Printf("Error fetching job posts: %v\n", err)
		return nil, fmt.Errorf("could not fetch job posts: %v", err)
	}

	// If no job posts are found, return an empty list and a message
	if len(jobPosts) == 0 {
		return nil, fmt.Errorf("no job posts available")
	}

	return jobPosts, nil
}