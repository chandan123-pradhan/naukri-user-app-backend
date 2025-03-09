package repositories

import (
	"fmt"
	"log"
	"strconv"

	"naurki_app_backend.com/config"
	"naurki_app_backend.com/models"
)

func GetJobPosts() ([]models.JobPost, error) {
	// Prepare the SQL query to fetch all job posts
	stmt := `SELECT id, jobTitle, jobDescription, qualification, noOfRequirements, contactEmail, contactNumber, jobLocation, skills, status, company_id, company_name, company_logo 
			 FROM job_post`

	// Execute the query and store the result
	rows, err := config.DB.Query(stmt)
	if err != nil {
		// Log the error and return an appropriate message
		fmt.Printf("Error querying job posts: %v\n", err)
		return nil, fmt.Errorf("could not query job posts: %v", err)
	}
	defer rows.Close()

	var jobPosts []models.JobPost
	for rows.Next() {
		var jobPost models.JobPost
		if err := rows.Scan(&jobPost.JobID, &jobPost.JobTitle, &jobPost.JobDescription, &jobPost.Qualification, &jobPost.NoOfRequirements, 
			&jobPost.ContactEmail, &jobPost.ContactNumber, &jobPost.JobLocation, &jobPost.Skills, &jobPost.Status, &jobPost.CompanyID, &jobPost.CompanyName, &jobPost.CompanyLogo); err != nil {
			// Log the error and return an appropriate message
			fmt.Printf("Error scanning job post: %v\n", err)
			return nil, fmt.Errorf("could not scan job post: %v", err)
		}
		if jobPost.Status == "open" {
			jobPosts = append(jobPosts, jobPost)
		}
	}

	// Check for any error that occurred during iteration
	if err := rows.Err(); err != nil {
		// Log the error and return an appropriate message
		fmt.Printf("Error iterating over job posts: %v\n", err)
		return nil, fmt.Errorf("error occurred while fetching job posts: %v", err)
	}

	return jobPosts, nil
}




func ApplyJob(userId int, jobId string) error {
	// Convert jobId to int (assuming jobId is stored as an integer in the database)
	jobIdInt, err := strconv.Atoi(jobId)
	if err != nil {
		return fmt.Errorf("invalid job ID format: %v", err)
	}

	// Check if the user has already applied for this job
	var count int
	stmt := `SELECT COUNT(*) FROM applications WHERE user_id = ? AND job_id = ?`
	err = config.DB.QueryRow(stmt, userId, jobIdInt).Scan(&count)
	if err != nil {
		return fmt.Errorf("failed to check application status: %v", err)
	}

	if count > 0 {
		// User has already applied for the job
		return fmt.Errorf("you have already applied this job")
	}

	// Insert the application record if not already applied
	stmt = `INSERT INTO applications (user_id, job_id, status) VALUES (?, ?, 'pending')`
	_, err = config.DB.Exec(stmt, userId, jobIdInt)
	if err != nil {
		log.Printf("Error applying for job: %v", err)  // Log the actual error
		return fmt.Errorf("failed to apply for the job: %v", err)  // Return the real error
	}

	// Return success if application was successful
	return nil
}


func GetJobDetailsWithApplicants(jobID, currentUserID string) (*models.JobPostDetails, error) {
	// Prepare the SQL query to fetch the job details based on job_id
	jobStmt := `SELECT id, jobTitle, jobDescription, qualification, noOfRequirements, contactEmail, contactNumber, jobLocation, skills, status, company_id, company_name, company_logo 
				FROM job_post WHERE id = ?`

	// Execute the query to get job details
	var jobPost models.JobPost
	err := config.DB.QueryRow(jobStmt, jobID).Scan(&jobPost.JobID, &jobPost.JobTitle, &jobPost.JobDescription, &jobPost.Qualification,
		&jobPost.NoOfRequirements, &jobPost.ContactEmail, &jobPost.ContactNumber, &jobPost.JobLocation, &jobPost.Skills,
		&jobPost.Status, &jobPost.CompanyID, &jobPost.CompanyName, &jobPost.CompanyLogo)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, fmt.Errorf("job not found")
		}
		fmt.Printf("Error fetching job details: %v\n", err)
		return nil, fmt.Errorf("could not fetch job details: %v", err)
	}

	// Prepare the SQL query to fetch the count of applicants who applied for this job (status = 'pending')
	applicantCountStmt := `SELECT COUNT(*) 
						   FROM applications a 
						   WHERE a.job_id = ? AND a.status = 'pending'`

	// Execute the query to get the number of applicants
	var applicantCount int
	err = config.DB.QueryRow(applicantCountStmt, jobID).Scan(&applicantCount)
	if err != nil {
		fmt.Printf("Error querying applicant count: %v\n", err)
		return nil, fmt.Errorf("could not query applicant count: %v", err)
	}

	// Prepare the SQL query to check if the current user has applied for this job
	isAppliedStmt := `SELECT COUNT(*) 
					  FROM applications a 
					  WHERE a.job_id = ? AND a.user_id = ?`

	// Execute the query to check if the current user has applied
	var isApplied int
	err = config.DB.QueryRow(isAppliedStmt, jobID, currentUserID).Scan(&isApplied)
	if err != nil {
		fmt.Printf("Error checking user application status: %v\n", err)
		return nil, fmt.Errorf("could not check application status: %v", err)
	}

	// Prepare the job details including the applicant count and applied status
	jobDetails := &models.JobPostDetails{
		JobPost:      jobPost,
		ApplicantCount: applicantCount,
		IsApplied:    isApplied > 0, // if count > 0, user has applied
	}

	return jobDetails, nil
}





func GetAppliedJobs(userID int) ([]models.AppliedJob, error) {
	fmt.Println("Fetching jobs for user:", userID) // Debug log

	stmt := `SELECT jp.id, jp.jobTitle, jp.jobDescription, jp.qualification, jp.noOfRequirements, 
	                jp.contactEmail, jp.contactNumber, jp.jobLocation, jp.skills, jp.status, 
	                jp.company_id, jp.company_name, jp.company_logo, a.application_date, a.status 
	         FROM job_post jp
	         JOIN applications a ON jp.id = a.job_id
	         WHERE a.user_id = ?;`

	
	rows, err := config.DB.Query(stmt, userID)
	if err != nil {
		fmt.Printf("Error querying applied jobs: %v\n", err)
		return nil, fmt.Errorf("could not query applied jobs: %v", err)
	}
	defer rows.Close()

	var appliedJobs []models.AppliedJob
	for rows.Next() {
		var job models.AppliedJob
		if err := rows.Scan(&job.JobID, &job.JobTitle, &job.JobDescription, &job.Qualification, 
			&job.NoOfRequirements, &job.ContactEmail, &job.ContactNumber, &job.JobLocation, 
			&job.Skills, &job.Status, &job.CompanyID, &job.CompanyName, &job.CompanyLogo, 
			&job.ApplicationDate, &job.ApplicationStatus); err != nil {
			fmt.Printf("Error scanning applied job: %v\n", err)
			return nil, fmt.Errorf("could not scan applied job: %v", err)
		}
		appliedJobs = append(appliedJobs, job)
	}

	if len(appliedJobs) == 0 {
		fmt.Println("No applied jobs found for user:", userID)
		return []models.AppliedJob{}, nil // Returning an empty list `[]`
	}

	// Check for any error that occurred during iteration
	if err := rows.Err(); err != nil {
		// Log the error and return an appropriate message
		fmt.Printf("Error iterating over applied jobs: %v\n", err)
		return nil, fmt.Errorf("error occurred while fetching applied jobs: %v", err)
	}

	return appliedJobs, nil
}


