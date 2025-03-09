package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"naurki_app_backend.com/config"
	"naurki_app_backend.com/routes"
)
//this is test

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set up database connection
	config.InitDB()

	// Initialize all routes
	router := routes.InitializeRoutes()

	// Get PORT from .env or set default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}

	log.Printf("Server started on :%s", port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, router)) // Bind to all interfaces
}
