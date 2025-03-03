package models

type JobPost struct {
	JobID          int    `json:"job_id"`
	JobTitle       string `json:"job_title"`
	JobDescription string `json:"job_description"`
	Qualification  string `json:"qualification"`
	NoOfRequirements int  `json:"no_of_requirements"`
	ContactEmail   string `json:"contact_email"`
	ContactNumber  string `json:"contact_number"`
	JobLocation    string `json:"job_location"`
	Skills         string `json:"skills"`
	Status         string `json:"status"`
	CompanyID      int    `json:"company_id"`
	CompanyLogo    string `json:"company_logo"`
	CompanyName    string `json:"company_name"`
}

// Applicant represents a user who applied for the job
type Applicant struct {
	UserID         int    `json:"user_id"`
	FullName       string `json:"full_name"`
	ProfileImageURL string `json:"profile_image_url"`
}

// JobPostDetails holds the job post along with the applicants
type JobPostDetails struct {
	JobPost       JobPost `json:"jobPost"`
	ApplicantCount int     `json:"applicantCount"`
	IsApplied     bool    `json:"isApplied"`
}