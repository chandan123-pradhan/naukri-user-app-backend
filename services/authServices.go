package services

import (
	"fmt"

	"math/rand"
	"naurki_app_backend.com/models"
	"naurki_app_backend.com/repositories"
	"naurki_app_backend.com/utils"
)


func RegisterUser(fullName, emailID, highestQualification, password, mobileNumber,
	preffered_location,
	preffered_sallary,
	preffered_shift,
	employment_type,
	description, 
	skills,
	profileImageURL string) (*models.User, error) {
	// Hash the password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	// Create a new user instance with the profile image URL
	user := &models.User{
		FullName:             fullName,
		EmailID:              emailID,
		HighestQualification: highestQualification,
		Password:             hashedPassword,
		MobileNumber:         mobileNumber,
		ProfileImageURL:      profileImageURL,
		PrefferedLocation:    preffered_location,
		PrefferedSallary:     preffered_sallary,
		PrefferedShift:       preffered_shift,
		Description:          description,
		EmploymentType:       employment_type,
		Skills:               skills,
	}

	// Save the user to the database (including profile image URL)
	if err := repositories.CreateUser(user); err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	return user, nil
}

// LoginUser handles user login, verifies credentials, and returns a user and JWT token
func LoginUser(emailID, password string) (*models.User, string, error) {
	// Validate email format
	if !utils.IsValidEmail(emailID) {
		return nil, "", fmt.Errorf("invalid email format")
	}

	// Fetch the user from the database by email ID
	user, err := repositories.GetUserByEmail(emailID)
	if err != nil {
		return nil, "", fmt.Errorf("user not found")
	}

	// Check if the provided password matches the stored password hash
	if valid := utils.CheckPasswordHash(password, user.Password); !valid {
		return nil, "", fmt.Errorf("incorrect password")
	}

	// Generate JWT token for the user
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return nil, "", fmt.Errorf("failed to generate token: %v", err)
	}

	// Return the sanitized user data (without the password) and the JWT token
	return &models.User{
		ID:                   user.ID,
		FullName:             user.FullName,
		EmailID:              user.EmailID,
		HighestQualification: user.HighestQualification,
		MobileNumber:         user.MobileNumber,
		ProfileImageURL:      user.ProfileImageURL,
		Skills: user.Skills,
		EmploymentType: user.EmploymentType,
		Description: user.Description,
		PrefferedLocation: user.PrefferedLocation,
		PrefferedSallary: user.PrefferedSallary,
		PrefferedShift: user.PrefferedShift,
	}, token, nil
}


func GenerateOTP(phone string) (string, error) {
	// Validate email format
	if !utils.IsValidPhoneNumber(phone){
		return "Invalid Phone Number",fmt.Errorf("invalid Phone Number")
	}
	otp := fmt.Sprintf("%04d", rand.Intn(10000))
	fmt.Println(otp)
	err :=repositories.GenerateOtp(phone,otp);
	// err :=repositories.GenerateOtp(phone,"1234");
	if(err!=nil){
		return "", fmt.Errorf("failed to save OTP: %v", err)
	}
	return "success",nil
}

func VerifyOtp(phone, otp string) (string, error) {
	// Validate email format
	if !utils.IsValidPhoneNumber(phone){
		return "Invalid Phone Number",fmt.Errorf("invalid Phone Number")
	}
	
	lastOneOtp,err := repositories.VerifyOtp(phone,otp)
	
	if(err!=nil){
		return "OTP Verification filed",err
	}

	if lastOneOtp == otp {
		return "success",nil
	}
	return "Incorrect OTP",fmt.Errorf("incorrect OTP")
	
}


func LoginViaOtp(phone, otp string)(*models.User, string, error){
	_, err:= VerifyOtp(phone,otp)
	if err !=nil{
		return nil,"",err
	}

	//if otp verified
	user, err := repositories.GetUserByMobileNumber(phone)
	if err != nil {
		return nil, "", fmt.Errorf("user not found")
	}
	// Generate JWT token for the user
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return nil, "", fmt.Errorf("failed to generate token: %v", err)
	}

	// Return the sanitized user data (without the password) and the JWT token
	return &models.User{
		ID:                   user.ID,
		FullName:             user.FullName,
		EmailID:              user.EmailID,
		HighestQualification: user.HighestQualification,
		MobileNumber:         user.MobileNumber,
		ProfileImageURL:      user.ProfileImageURL,
		Skills: user.Skills,
		EmploymentType: user.EmploymentType,
		Description: user.Description,
		PrefferedLocation: user.PrefferedLocation,
		PrefferedSallary: user.PrefferedSallary,
		PrefferedShift: user.PrefferedShift,
	}, token, nil

}



