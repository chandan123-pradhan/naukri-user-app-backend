package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"naurki_app_backend.com/services"
	"naurki_app_backend.com/utils"
)

// Register handles user registration, including profile picture upload and token generation
func Register(w http.ResponseWriter, r *http.Request) {
	// Parse multipart form data
	err := r.ParseMultipartForm(10 << 20) // 10MB limit for file size
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, "Error parsing request", map[string]interface{}{})
		return
	}

	// Extract user information from form data
	var req struct {
		FullName             string `json:"full_name"`
		EmailID              string `json:"email_id"`
		HighestQualification string `json:"highest_qualification"`
		Password             string `json:"password"`
		ConfirmPassword      string `json:"confirm_password"`
		MobileNumber         string `json:"mobile_number"`
		PrefferedLocation    string `json:"preffered_location"`
		PrefferedSallary     string `json:"preffered_sallary"`
		PrefferedShift       string `json:"preffered_shift"`
		EmploymentType       string `json:"employment_type"`
		Description          string `json:"description"`
		Skills               string `json:"skills"`
	}

	// Retrieve fields from multipart form
	req.FullName = r.FormValue("full_name")
	req.EmailID = r.FormValue("email_id")
	req.HighestQualification = r.FormValue("highest_qualification")
	req.Password = r.FormValue("password")
	req.ConfirmPassword = r.FormValue("confirm_password")
	req.MobileNumber = r.FormValue("mobile_number")
	req.PrefferedLocation=r.FormValue("preffered_location")
	req.PrefferedSallary=r.FormValue("preffered_sallary")
	req.PrefferedShift=r.FormValue("preffered_shift")
	req.EmploymentType=r.FormValue("employment_type")
	req.Description=r.FormValue("description")
	req.Skills=r.FormValue("skills")

	// Validate input
	if err := utils.ValidateRegistrationInput(req); err != nil {
		respondWithJSON(w, http.StatusBadRequest, err.Error(), map[string]interface{}{})
		return
	}

	// Check for profile image in the form
	var profileImageURL string
	file, _, err := r.FormFile("profile_image")
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, "Profile image is required", map[string]interface{}{})
		return
	}
	defer file.Close()

	// Save the profile image to the server (or use cloud storage)
	// Generate a unique file name based on timestamp
	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), "profile_image.jpg")
	filePath := filepath.Join("./uploads", fileName)

	// Create the file on the server
	outFile, err := os.Create(filePath)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, "Error saving file", map[string]interface{}{})
		return
	}
	defer outFile.Close()

	// Copy the content from the uploaded file to the server file
	_, err = outFile.ReadFrom(file)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, "Error writing file", map[string]interface{}{})
		return
	}

	// Now that the file is uploaded, save the file path (URL) to the database
	profileImageURL = "/uploads/" + fileName // Adjust this based on where your images are served

	// Call service to register the user with profile image URL
	user, err := services.RegisterUser(req.FullName, req.EmailID, req.HighestQualification, req.Password, req.MobileNumber,
		req.PrefferedLocation,
		req.PrefferedSallary,
		req.PrefferedShift,
		req.EmploymentType,
		req.Description, 
		req.Skills,
		profileImageURL)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, "Email ID already Exists, try to login please", map[string]interface{}{})
		return
	}

	// Generate JWT token for the user
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, "Failed to generate token", map[string]interface{}{})
		return
	}

	// Prepare a sanitized user response (without password)
	userResponse := map[string]interface{}{
		"id":                   user.ID,
		"full_name":            user.FullName,
		"email_id":             user.EmailID,
		"highest_qualification": user.HighestQualification,
		"mobile_number":        user.MobileNumber,
		"profile_image_url":    user.ProfileImageURL,
		"skills":				user.Skills,
		"preffered_location":	user.PrefferedLocation,
		"preffered_sallary":	user.PrefferedSallary,
		"preffered_shift":user.PrefferedShift,
		"description":user.Description,
		"employment_type":user.EmploymentType,
	}

	// Respond with success and include the JWT token
	respondWithJSON(w, http.StatusCreated, "Registration successful", map[string]interface{}{
		"user":  userResponse, // User data without password
		"token": token,        // The JWT token
	})
}











func Login(w http.ResponseWriter, r *http.Request) {
	// Parse the incoming request to get login credentials
	var req struct {
		EmailID  string `json:"email_id"`
		Password string `json:"password"`
	}

	// Decode the incoming JSON request into the req object
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, "Invalid request format", map[string]interface{}{})
		return
	}

	// Validate the input fields
	if req.EmailID == "" || req.Password == "" {
		respondWithJSON(w, http.StatusBadRequest, "Email and Password are required", map[string]interface{}{})
		return
	}

	// Call service to authenticate the user
	user, token, err := services.LoginUser(req.EmailID, req.Password)
	if err != nil {
		respondWithJSON(w, http.StatusUnauthorized,"Email and Passwore are not mached", map[string]interface{}{})
		return
	}

	// Prepare the response user data (without password)
	userResponse := map[string]interface{}{
		"id":                   user.ID,
		"full_name":            user.FullName,
		"email_id":             user.EmailID,
		"highest_qualification": user.HighestQualification,
		"mobile_number":        user.MobileNumber,
		"profile_image_url":    user.ProfileImageURL,
		"description":user.Description,
		"skills":user.Skills,
		"employment_type":user.EmploymentType,
		"preffered_location":user.PrefferedLocation,
		"preffered_sallary":user.PrefferedSallary,
		"preffered_shift":user.PrefferedShift,

	}

	// Respond with success and include the JWT token
	respondWithJSON(w, http.StatusOK, "Login successful", map[string]interface{}{
		"user":  userResponse, // User data without password
		"token": token,        // The JWT token
	})
}


// GenerateOTP generates and sends OTP to the user
func GenerateOTP(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Phone  string `json:"phone"`
		
	}
	
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&req); err != nil {
		respondWithJSON(w, http.StatusInternalServerError, "Invalid request format", map[string]interface{}{})
		return
	}
	if req.Phone == "" {
		respondWithJSON(w, http.StatusBadRequest, "Phone is required fields", map[string]interface{}{})
		return
	}
	_, err := services.GenerateOTP(req.Phone)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, err.Error(), map[string]interface{}{})
		return
	}

	respondWithJSON(w, http.StatusOK, "OTP Generate Successfully Done.", map[string]interface{}{})
}




func VerifyOtp(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Phone  string `json:"phone"`
		Otp	   string `json:"otp"`
		
	}
	
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&req); err != nil {
		respondWithJSON(w, http.StatusInternalServerError,"Invalid request format", map[string]interface{}{})
		return
	}
	if req.Phone == "" || req.Otp == "" {
		respondWithJSON(w, http.StatusBadRequest, "Phone and Otp are required fields", map[string]interface{}{})
		return
	}
	_,err := services.VerifyOtp(req.Phone,req.Otp);
	if(err!=nil){
		respondWithJSON(w, http.StatusBadRequest, err.Error(), map[string]interface{}{})
		return
	}

	respondWithJSON(w, http.StatusOK, "OTP Verified Successfully Done.", map[string]interface{}{})
}







func LoginViaOtp(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Phone  string `json:"phone"`
		Otp	   string `json:"otp"`
		
	}
	
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&req); err != nil {
		respondWithJSON(w, http.StatusInternalServerError, "Invalid request format", map[string]interface{}{})
		return
	}
	if req.Phone == "" || req.Otp == "" {
		respondWithJSON(w, http.StatusBadRequest, "Phone and Otp are required fields", map[string]interface{}{})
		return
	}

	user, token, err := services.LoginViaOtp(req.Phone,req.Otp)
	if err != nil {
		respondWithJSON(w, http.StatusUnauthorized,err.Error(), map[string]interface{}{})
		return
	}

	// Prepare the response user data (without password)
	userResponse := map[string]interface{}{
		"id":                   user.ID,
		"full_name":            user.FullName,
		"email_id":             user.EmailID,
		"highest_qualification": user.HighestQualification,
		"mobile_number":        user.MobileNumber,
		"profile_image_url":    user.ProfileImageURL,
		"description":user.Description,
		"skills":user.Skills,
		"employment_type":user.EmploymentType,
		"preffered_location":user.PrefferedLocation,
		"preffered_sallary":user.PrefferedSallary,
		"preffered_shift":user.PrefferedShift,

	}

	// Respond with success and include the JWT token
	respondWithJSON(w, http.StatusOK, "Login successful", map[string]interface{}{
		"user":  userResponse, // User data without password
		"token": token,        // The JWT token
	})

}





// Helper function to respond with JSON in the desired structure
func respondWithJSON(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// Set the response structure
	response := map[string]interface{}{
		"status":  "failure", // Default status is failure
		"message": message,
		"data":    data, // Send the provided data (or empty if nil)
	}

	// If there's no error, set the status to "success"
	if statusCode == http.StatusOK || statusCode == http.StatusCreated {
		response["status"] = "success"
	}

	// Encode the response as JSON
	json.NewEncoder(w).Encode(response)
}

