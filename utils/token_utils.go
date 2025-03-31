package utils

import (
	"fmt"
	"net/http"
	"strings"
)

// Helper function for validating the JWT token and extracting the user ID
func ValidateToken(w http.ResponseWriter, r *http.Request) (int, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		RespondWithJSON(w, http.StatusUnauthorized, "Authorization token is required", nil)
		return 0, fmt.Errorf("authorization token is required")
	}

	// Extract the Bearer token
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		RespondWithJSON(w, http.StatusUnauthorized, "Invalid token format", nil)
		return 0, fmt.Errorf("invalid token format")
	}

	// Validate the token and get the user ID
	userID, err := VerifyJWT(tokenString)
	if err != nil {
		RespondWithJSON(w, http.StatusUnauthorized, err.Error(), nil)
		return 0, err
	}

	return userID, nil
}