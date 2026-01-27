# Simple Go Web Server

A basic web server built with Go's standard library that demonstrates routing, handlers, middleware, and JSON APIs.

## Features

- **Multiple Routes**: Home, Hello, Time, and API endpoints
- **Query Parameters**: Dynamic greeting with URL parameters
- **JSON API**: RESTful endpoint returning JSON data
- **Logging Middleware**: Automatic request logging with timing
- **Proper Configuration**: Timeout settings and graceful server setup

## Prerequisites

- Go 1.16 or higher installed on your system
- Basic knowledge of terminal/command line

## Installation & Running

1. **Initialize the Go module** (if not already done):
   ```bash
   go mod init simple-webserver
   ```

2. **Run the server**:
   ```bash
   go run main.go
   ```

3. **Build an executable** (optional):
   ```bash
   go build -o webserver main.go
   ./webserver
   ```

## Available Endpoints

Once the server is running on `http://localhost:8080`, you can access:

- **`/`** - Home page with links to all endpoints
- **`/hello`** - Simple greeting page
- **`/hello?name=YourName`** - Personalized greeting
- **`/time`** - Displays current server time
- **`/api/data`** - JSON API endpoint

## Example Usage

### Using curl:

```bash
# Get home page
curl http://localhost:8080/

# Get JSON data
curl http://localhost:8080/api/data

# Personalized greeting
curl http://localhost:8080/hello?name=Alice
```

### Using a browser:

Simply open `http://localhost:8080` in your browser and navigate through the links.

## Code Explanation

### Key Components:

1. **Handler Functions**: Each route has a handler function that processes requests
2. **ServeMux**: Go's built-in router that maps URLs to handlers
3. **Middleware**: The `loggingMiddleware` wraps handlers to add logging
4. **Server Configuration**: Timeout settings for production-ready setup

### Customization:

- Change the port by modifying the `port` variable in `main()`
- Add new routes by creating handler functions and registering them with `mux.HandleFunc()`
- Modify timeout values in the `http.Server` configuration

## Next Steps

To expand this server, you could add:
- Static file serving
- Database connections
- Template rendering
- More complex routing (using packages like gorilla/mux)
- Authentication/authorization
- CORS middleware
- POST request handling
- Form processing

## Stopping the Server

Press `Ctrl+C` in the terminal to gracefully stop the server.
