package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// User represents a user structure
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// In-memory storage (for demo purposes)
var users []User

// Handler for form page
func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Display the form
		html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>User Form</title>
			<style>
				body { font-family: Arial, sans-serif; max-width: 600px; margin: 50px auto; padding: 20px; }
				input, button { padding: 10px; margin: 5px 0; width: 100%; box-sizing: border-box; }
				button { background: #007bff; color: white; border: none; cursor: pointer; }
				button:hover { background: #0056b3; }
			</style>
		</head>
		<body>
			<h1>Submit User Information</h1>
			<form method="POST" action="/form">
				<input type="text" name="name" placeholder="Name" required>
				<input type="email" name="email" placeholder="Email" required>
				<button type="submit">Submit</button>
			</form>
			<p><a href="/users">View All Users</a></p>
		</body>
		</html>
		`
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, html)
		return
	}

	if r.Method == http.MethodPost {
		// Parse form data
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}

		name := r.FormValue("name")
		email := r.FormValue("email")

		// Add user to storage
		user := User{Name: name, Email: email}
		users = append(users, user)

		log.Printf("New user added: %s (%s)", name, email)

		// Redirect to users list
		http.Redirect(w, r, "/users", http.StatusSeeOther)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

// Handler to display all users
func usersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Users List</title>
		<style>
			body { font-family: Arial, sans-serif; max-width: 600px; margin: 50px auto; padding: 20px; }
			table { width: 100%; border-collapse: collapse; margin: 20px 0; }
			th, td { border: 1px solid #ddd; padding: 12px; text-align: left; }
			th { background: #007bff; color: white; }
			tr:nth-child(even) { background: #f2f2f2; }
		</style>
	</head>
	<body>
		<h1>All Users</h1>
	`)

	if len(users) == 0 {
		fmt.Fprint(w, "<p>No users yet. <a href='/form'>Add one!</a></p>")
	} else {
		fmt.Fprint(w, "<table><tr><th>Name</th><th>Email</th></tr>")
		for _, user := range users {
			fmt.Fprintf(w, "<tr><td>%s</td><td>%s</td></tr>", user.Name, user.Email)
		}
		fmt.Fprint(w, "</table>")
	}

	fmt.Fprint(w, "<p><a href='/form'>Add New User</a> | <a href='/'>Home</a></p>")
	fmt.Fprint(w, "</body></html>")
}

// API handler to get users as JSON
func apiUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	if r.Method == http.MethodGet {
		json.NewEncoder(w).Encode(users)
		return
	}

	if r.Method == http.MethodPost {
		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		users = append(users, user)
		log.Printf("New user added via API: %s (%s)", user.Name, user.Email)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

// Home page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	html := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Advanced Go Web Server</title>
		<style>
			body { font-family: Arial, sans-serif; max-width: 800px; margin: 50px auto; padding: 20px; }
			.card { border: 1px solid #ddd; padding: 20px; margin: 10px 0; border-radius: 5px; }
			.card:hover { box-shadow: 0 4px 8px rgba(0,0,0,0.1); }
			h1 { color: #007bff; }
			a { color: #007bff; text-decoration: none; }
			a:hover { text-decoration: underline; }
		</style>
	</head>
	<body>
		<h1>🚀 Advanced Go Web Server</h1>
		
		<div class="card">
			<h2>📝 Form Handling</h2>
			<p><a href="/form">Submit a form</a> to add users</p>
		</div>

		<div class="card">
			<h2>👥 User Management</h2>
			<p><a href="/users">View all users</a> in HTML format</p>
		</div>

		<div class="card">
			<h2>🔌 REST API</h2>
			<p>GET/POST to <code>/api/users</code> for JSON operations</p>
			<p>Example: <code>curl http://localhost:8080/api/users</code></p>
		</div>
	</body>
	</html>
	`
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, html)
}

func main() {
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/form", formHandler)
	mux.HandleFunc("/users", usersHandler)
	mux.HandleFunc("/api/users", apiUsersHandler)

	// Server configuration
	port := ":8080"
	server := &http.Server{
		Addr:         port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Printf("🚀 Server starting on http://localhost%s", port)
	log.Println("Available endpoints:")
	log.Println("  - http://localhost:8080/")
	log.Println("  - http://localhost:8080/form")
	log.Println("  - http://localhost:8080/users")
	log.Println("  - http://localhost:8080/api/users")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
