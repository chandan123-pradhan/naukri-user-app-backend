package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/cors" // ✨ Import CORS

	"naurki_app_backend.com/config"
	"naurki_app_backend.com/firebase"
	"naurki_app_backend.com/routes"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set up database connection
	config.InitDB()
	firebase.InitFirebase()

	// Initialize all routes
	router := routes.InitializeRoutes()

	// Set up CORS middleware
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // 🔒 Replace "*" with specific origin(s) in production
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}).Handler(router)

	// Get PORT from .env or set default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}

	log.Printf("🚀 Server started on :%s", port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, corsHandler)) // 👈 Wrap with corsHandler
}
