package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"naurki_app_backend.com/models"
)

// Secret key for signing JWTs (in a production app, store it in an environment variable)
var jwtSecret = []byte("your-secret-key")
func GenerateJWT(userID int) (string, error) {
	// Create a new JWT token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims (payload) for the token
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = userID    // Ensure "id" is set with the user ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %v", err)
	}

	return tokenString, nil
}

func VerifyJWT(tokenString string) (int, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token's signing method matches
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return 0, fmt.Errorf("token is expired")
	}

	// Ensure the token is valid and extract the claims
	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
		// Print the claims for debugging purposes
		fmt.Printf("Claims: %+v\n", claims)

		// If valid, return the user ID from the claims
		return claims.UserID, nil
	}

	return 0, fmt.Errorf("invalid token")
}
