package models

type SendNotificationRequest struct {
	FcmToken string `json:"fcm_token"`
	Title    string `json:"title"`
	Body     string `json:"body"`
}
