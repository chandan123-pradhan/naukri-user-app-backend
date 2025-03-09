package models

type AppliedJob struct {
	JobID            int
	JobTitle         string
	JobDescription   string
	Qualification    string
	NoOfRequirements int
	ContactEmail     string
	ContactNumber    string
	JobLocation      string
	Skills           string
	Status           string
	CompanyID        int
	CompanyName      string
	CompanyLogo      string
	ApplicationDate  string
	ApplicationStatus string // 'pending', 'accepted', 'rejected'
}
