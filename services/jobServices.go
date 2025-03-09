package services

import (
	"fmt"
	"strconv"

	"naurki_app_backend.com/models"
	"naurki_app_backend.com/repositories"
)

// GetJobPosts fetches all job posts from the repository
func GetJobPosts() ([]models.JobPost, error) {
	// Fetch the job posts from the repository
	jobPosts, err := repositories.GetJobPosts()
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

func IsJobIdCorrect(jobId string) bool {
    // Get job posts (assuming GetJobPosts() returns a list of job posts)
    jobPosts, _ := GetJobPosts()

    // Convert jobId string to integer
    jobIdInt, err := strconv.Atoi(jobId)
    if err != nil {
        fmt.Println("Invalid jobId format:", err)
        return false
    }

    // Loop through job posts and check for match
    for i := 0; i < len(jobPosts); i++ {
        if jobPosts[i].JobID == jobIdInt {
            return true // Job ID matched
        }
    }
    
    return false // Job ID not found
}

func ApplyJob(userID int ,jobId string)error{
	
	applicationMessage := repositories.ApplyJob(userID, jobId)

	// You can do additional logic after calling ApplyJob if needed
	// For now, just return the message received from ApplyJob
	return applicationMessage
	 
}


// JobDetails fetches the job details along with the applicants
func JobDetails(jobId string, currentUserId string) (*models.JobPostDetails, error) {
	// Call the repository function to get job details and applicants
	jobDetails, err := repositories.GetJobDetailsWithApplicants(jobId,currentUserId)
	if err != nil {
		// Handle error, return a meaningful message
		return nil, fmt.Errorf("could not fetch job details: %v", err)
	}

	// Optionally, you can add more business logic here if needed

	// Return the job details
	return jobDetails, nil
}

func GetAppliedJobs(userId int) ([]models.AppliedJob, error) {
	jobPosts, err := repositories.GetAppliedJobs(userId)
	if err != nil {
		// Log the error and return a more detailed error message
		fmt.Printf("Error fetching applied jobs: %v\n", err)
		return nil, fmt.Errorf("could not fetch applied jobs: %v", err)
	}

	// ✅ FIX: Return an empty list (`[]`) instead of an error if no applied jobs are found
	if len(jobPosts) == 0 {
		fmt.Println("No applied jobs found for user:", userId)
		return []models.AppliedJob{}, nil // Returning an empty list `[]` instead of an error
	}

	return jobPosts, nil
}
