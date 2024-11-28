package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Database configuration
const (
	dsn             = "root:24680@tcp(127.0.0.1:3306)/time_logger" // DSN
	torontoTimeZone = "America/Toronto"                            // Time zone
)

var db *sql.DB

// Response structure
type Response struct {
	CurrentTime string `json:"current_time"`
	Timezone    string `json:"timezone"`
}

func main() {
	var err error

	// Connect to the database
	db, err = connectToDatabase()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	// Register the /current-time endpoint
	http.HandleFunc("/current-time", currentTimeHandler)

	// Start the HTTP server
	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func connectToDatabase() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("database connection failed: %w", err)
	}

	log.Println("Connected to MySQL database successfully!")
	return db, nil
}

func currentTimeHandler(w http.ResponseWriter, r *http.Request) {
	// Load Toronto timezone
	location, err := time.LoadLocation(torontoTimeZone)
	if err != nil {
		http.Error(w, "Failed to load timezone: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Get current time in Toronto timezone
	currentTime := time.Now().In(location)

	// Insert the time into the database
	if err := logTimeToDatabase(currentTime); err != nil {
		http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the current time as JSON response
	response := Response{
		CurrentTime: currentTime.Format("2006-01-02 15:04:05"),
		Timezone:    torontoTimeZone,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
	}
}

func logTimeToDatabase(timestamp time.Time) error {
	query := "INSERT INTO time_log (timestamp) VALUES (?)"
	_, err := db.Exec(query, timestamp)
	if err != nil {
		log.Printf("Error logging time to database: %v", err)
		return fmt.Errorf("failed to insert time into database: %w", err)
	}

	log.Printf("Time logged to database: %s", timestamp.Format("2006-01-02 15:04:05"))
	return nil
}
