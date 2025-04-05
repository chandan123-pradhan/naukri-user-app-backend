package services

import (
	"context"
	"fmt"
	"log"

	"naurki_app_backend.com/firebase"
	"naurki_app_backend.com/repositories"

	firebase_messaging "firebase.google.com/go/v4/messaging"
)

func SendNotificationToToken(token string, title string, body string) error {
    ctx := context.Background()

    client := firebase.GetClient() // gets the initialized Firebase client

    message := &firebase_messaging.Message{
        Token: token,
        Notification: &firebase_messaging.Notification{
            Title: title,
            Body:  body,
        },
    }

    response, err := client.Send(ctx, message)
    if err != nil {
        log.Printf("Failed to send notification: %v", err)
        return err
    }

    log.Printf("Successfully sent message: %s", response)
    return nil
}



func GetCompanyFcm(companyId int) (string, error) {
	// Call the repository function to get the FCM token
	fcmToken, err := repositories.GetCompanyFCM(companyId)
	if err != nil {
		// Return an error if the token could not be retrieved
		return "", fmt.Errorf("failed to retrieve FCM token for company ID %d: %v", companyId, err)
	}
	
	// Return the FCM token if found successfully
	return fcmToken, nil
}