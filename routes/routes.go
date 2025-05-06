package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"naurki_app_backend.com/controllers"
)

func InitializeRoutes() *mux.Router {
	// Initialize the router
	router := mux.NewRouter()

	// Authentication Routes
	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/update_employment",controllers.UpdateEmploymentHistory).Methods("POST")
	router.HandleFunc("/get_profile",controllers.GetUserDetails).Methods("GET")
	router.HandleFunc("/get_jobs",controllers.GetJobPosts).Methods("GET")
	router.HandleFunc("/apply_job",controllers.ApplyJob).Methods("POST")
	router.HandleFunc("/job_details",controllers.GetJobDetails).Methods("POST")
	router.HandleFunc("/applied-jobs",controllers.GetAppliedJobs).Methods("GET")
	router.HandleFunc("/generate-otp",controllers.GenerateOTP).Methods("POST")
	router.HandleFunc("/verify-otp",controllers.VerifyOtp).Methods("POST")
	router.HandleFunc("/login-via-otp",controllers.LoginViaOtp).Methods("POST")
	router.HandleFunc("/get_jobs",controllers.SearchJobByTitle).Methods("POST")
	router.HandleFunc("/create-alert",controllers.CreateAlerts).Methods("POST")
	router.HandleFunc("/get-alert",controllers.GetAlerts).Methods("GET")
	router.HandleFunc("/send-notification",controllers.SendNotification).Methods("POST")
	router.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	// Add other routes for authentication or any other resources
	// For example:
	// router.HandleFunc("/login", controllers.Login).Methods("POST")

	// You can add more routes here as your application grows
	// Example: router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	// Example: router.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET")

	return router
}
