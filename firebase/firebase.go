package firebase

import (
    "context"
    "log"

    firebase "firebase.google.com/go/v4"
    "firebase.google.com/go/v4/messaging"
    "google.golang.org/api/option"
)

var client *messaging.Client // unexported

// GetClient returns the messaging client
func GetClient() *messaging.Client {
    return client
}

// InitFirebase initializes Firebase and sets up messaging client
func InitFirebase() {
    ctx := context.Background()

    opt := option.WithCredentialsFile("firebase/serviceAccountKey.json")
    app, err := firebase.NewApp(ctx, nil, opt)
    if err != nil {
        log.Fatalf("error initializing app: %v", err)
    }

    client, err = app.Messaging(ctx)
    if err != nil {
        log.Fatalf("error initializing messaging client: %v", err)
    }

    log.Println("Firebase initialized successfully.")
}
