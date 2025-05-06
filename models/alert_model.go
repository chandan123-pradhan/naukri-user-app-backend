package models


type CreateAlertModel struct {
	JobTitle string `json:"job_title" validate:"required"`
	Skills      string `json:"skills" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	UserName    string `json:"user_name" validate:"required"`
	MobileNo    string `json:"mobile_number" validate:"required,len=10"`
	Location    string `json:"location" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type UserAlert struct {
	ID              int    `json:"id"`
	JobTitle        string `json:"job_title"`
	Skills          string `json:"skills"`
	Email           string `json:"email"`
	UserName        string `json:"user_name"`
	MobileNo        string `json:"mobile_number"`
	UserID          int    `json:"user_id"`
	Location        string `json:"location"`
	Description     string `json:"description"`
	ProfileImageUrl string `json:"profile_image_url"`
}
