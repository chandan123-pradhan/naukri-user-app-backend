package controllers

import (
	"net/http"
	"path/filepath"
	"os"
)

func ServeImage(w http.ResponseWriter, r *http.Request) {
	// Extract the image filename from the URL path
	imageName := r.URL.Path[len("/uploads/"):] // Removes "/uploads/" prefix
	if imageName == "" {
		http.Error(w, "Image not found", http.StatusNotFound)
		return
	}

	// Define the path to the image file
	filePath := filepath.Join("./uploads", imageName)

	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, "Image not found", http.StatusNotFound)
		return
	}

	// Serve the image file
	w.Header().Set("Content-Type", "image/jpeg") // Assuming JPEG images; adjust as needed
	http.ServeFile(w, r, filePath)
}