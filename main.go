package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Handler for the home page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to the Go Web Server!</h1>")
	fmt.Fprintf(w, "<p>Available endpoints:</p>")
	fmt.Fprintf(w, "<ul>")
	fmt.Fprintf(w, "<li><a href='/'>Home</a></li>")
	fmt.Fprintf(w, "<li><a href='/hello'>Hello</a></li>")
	fmt.Fprintf(w, "<li><a href='/time'>Current Time</a></li>")
	fmt.Fprintf(w, "<li><a href='/api/data'>API Data (JSON)</a></li>")
	fmt.Fprintf(w, "</ul>")
}

// Handler for the hello page
func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}
	fmt.Fprintf(w, "<h1>Hello, %s!</h1>", name)
	fmt.Fprintf(w, "<p>Try: <a href='/hello?name=YourName'>/hello?name=YourName</a></p>")
}

// Handler for the time page
func timeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Fprintf(w, "<h1>Current Server Time</h1>")
	fmt.Fprintf(w, "<p>%s</p>", currentTime)
}

// Handler for JSON API endpoint
func apiDataHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")
	
	// Create sample data
	data := map[string]interface{}{
		"message": "Hello from the API!",
		"timestamp": time.Now().Unix(),
		"server": "Go Web Server",
		"version": "1.0.0",
	}
	
	// Encode and send JSON response
	json.NewEncoder(w).Encode(data)
}

// Middleware to log requests
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		
		next(w, r)
		
		log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	}
}

func main() {
	// Create a new ServeMux (router)
	mux := http.NewServeMux()
	
	// Register routes with middleware
	mux.HandleFunc("/", loggingMiddleware(homeHandler))
	mux.HandleFunc("/hello", loggingMiddleware(helloHandler))
	mux.HandleFunc("/time", loggingMiddleware(timeHandler))
	mux.HandleFunc("/api/data", loggingMiddleware(apiDataHandler))
	
	// Server configuration
	port := ":8080"
	server := &http.Server{
		Addr:         port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	
	log.Printf("Starting server on http://localhost%s", port)
	log.Printf("Press Ctrl+C to stop the server")
	
	// Start the server
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
