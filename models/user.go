package models

// User represents a user in the system
type User struct {
	ID                   int    `json:"id"`
	FullName             string `json:"full_name"`
	EmailID              string `json:"email_id"`
	HighestQualification string `json:"highest_qualification"`
	Password             string `json:"password"`
	MobileNumber         string `json:"mobile_number"`
	ProfileImageURL      string `json:"profile_image_url"` // New field for storing profile image URL
	CreatedAt            string `json:"created_at"`
	UpdatedAt            string `json:"updated_at"`
	PrefferedLocation    string `json:"preffered_location"`
	PrefferedSallary     string `json:"preffered_sallary"`
	PrefferedShift       string `json:"preffered_shift"`
	EmploymentType       string `json:"employment_type"`
	Description          string `json:"description"`
	Skills               string `json:"skills"`
}
