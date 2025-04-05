package repositories

import (
	"database/sql"
	"fmt"

	"naurki_app_backend.com/config" // Ensure this is where your models are located
)

// GetCompanyFCM will fetch the FCM token for a given user ID (or company ID)
func GetCompanyFCM(companyId int) (string, error) {
	fmt.Println("Fetching FCM token for user:", companyId) // Debug log

	// Prepare SQL query to fetch the FCM token for the company
	stmt := `SELECT fcm_token
			 FROM company_fcm_tokens
			 WHERE company_id = ?;`

	var fcmToken string
	err := config.DB.QueryRow(stmt, companyId).Scan(&fcmToken)
	if err != nil {
		// Handling no rows error
		if err == sql.ErrNoRows {
			fmt.Printf("No FCM token found for user ID %d\n", companyId)
			return "", fmt.Errorf("no FCM token found for user ID %d", companyId)
		}
		// Other errors (e.g., query errors)
		fmt.Printf("Error querying FCM token for user ID %d: %v\n", companyId, err)
		return "", fmt.Errorf("could not query FCM token: %v", err)
	}

	// Return the FCM token if it was found
	fmt.Printf("Successfully fetched FCM token for user ID %d\n", companyId)
	return fcmToken, nil
}
