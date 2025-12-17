# Getting Started with Go HTTP Programming – A Beginner's Guide

## 1. Title & Objective

**What technology did you choose?**  
Go (Golang) HTTP programming

**Why did you choose it?**
Go is a modern, statically-typed programming language designed for simplicity and efficiency. I chose Go HTTP programming because it provides built-in libraries for creating web servers with minimal boilerplate code, making it perfect for beginners who want to understand backend development fundamentals.

Additionally, Go powers many essential DevOps tools including:
- Docker (containerization)
- Kubernetes (container orchestration)
- Terraform (infrastructure as code)
- Prometheus (monitoring)
- Jenkins-X (CI/CD)
- Grafana (visualization)
- ArgoCD (GitOps)
- Istio (service mesh)

This makes Go skills highly valuable in DevOps and cloud-native development roles.

**What's my  end goal?**  
Create a simple HTTP web server that responds to HTTP requests and demonstrates core Go programming concepts including packages, functions, and the net/http library.

## 2. Quick Summary of the Technology

Go (Golang) is a compiled programming language developed by Google that emphasizes simplicity, concurrency, and performance. The net/http package provides built-in support for HTTP servers and clients.

**What is it?**  
A programming language with native HTTP capabilities for building web servers and applications.

**Where is it used?**  
Backend services, APIs, microservices, and web applications (used by companies like Uber, Netflix, and Dropbox).

**One real-world example:**  
Netflix uses Go for their edge services and content delivery networks due to its efficiency in handling concurrent connections.

## 3. System Requirements

**OS:** Linux/Mac/Windows  
**Tools/Editors required:** VS Code, GoLand, or any text editor  
**Go Version:** Go 1.19 or later  
**Packages:** Standard library only (no external dependencies required)

## 4. Installation & Setup Instructions

### Step 1: Install Go

**Windows:**
1. Download installer from https://golang.org/dl/
2. Run the MSI installer
3. Verify installation:
```bash
go version
```

**macOS:**
```bash
# Using Homebrew
brew install go

# Verify installation
go version
```

**Linux (Ubuntu/Debian):**
```bash
sudo apt update
sudo apt install golang-go

# Verify installation
go version
```

### Step 2: Set Up Your Workspace

1. Create a project directory:
```bash
mkdir go-http-demo
cd go-http-demo
```

2. Initialize Go module:
```bash
go mod init go-http-demo
```

### Step 3: Create Your First HTTP Server

Create a file called `main.go` with the following content:

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Register a handler function for the root path
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
    })

    fmt.Println("Server starting on http://localhost:8080")
    
    // Start the server on port 8080
    http.ListenAndServe(":8080", nil)
}
```

### Step 4: Run Your Server

```bash
go run main.go
```

You should see: "Server starting on http://localhost:8080"

### Step 5: Test Your Server

Open your browser and visit: http://localhost:8080

You should see: "Hello, World!"

## 5. Minimal Working Example

**What the example does:**  
Creates a simple HTTP server that listens on port 8080 and responds to all HTTP requests with "Hello, World!"

**Expected output:**  
- Terminal: "Server starting on http://localhost:8080"
- Browser: "Hello, World!"

**Code with inline comments:**

```go
package main

import (
    "fmt"        // For formatted I/O (printing)
    "net/http"   // HTTP server and client functionality
)

func main() {
    // Register a handler function for the root path "/"
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Write "Hello, World!" to the response
        fmt.Fprintf(w, "Hello, World!")
    })

    // Print startup message to console
    fmt.Println("Server starting on http://localhost:8080")
    
    // Start HTTP server on port 8080, nil uses default router
    http.ListenAndServe(":8080", nil)
}
```

**Key Concepts Explained:**

1. **package main**: Defines this as an executable program
2. **import**: Imports necessary libraries
3. **http.HandleFunc**: Registers a function to handle HTTP requests
4. **http.ResponseWriter**: Interface for writing HTTP responses
5. **http.Request**: Structure containing request details
6. **http.ListenAndServe**: Starts the HTTP server

## 6. AI Prompt Journal

### Prompt 1: Basic Go HTTP Server Setup
**Prompt I used:**  
"Give me a simple Go program that creates an HTTP server and responds with 'Hello, World!'"

**AI's response summary:**  
The AI provided a complete, working example using the net/http package with proper package structure and error handling.

**Brief part of the response that addresses the problem:**  
```go
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
})
```

**My evaluation of its helpfulness:**  
Very helpful - provided exactly what was needed with clear, idiomatic Go code. The response included proper imports and the fundamental pattern for creating HTTP handlers.

### Prompt 2: Understanding Go HTTP Concepts
**Prompt used:**  
"Explain how the net/http package works in Go and what http.HandleFunc does"

**AI's response summary:**  
The AI explained the HTTP handler pattern, how the ResponseWriter and Request interfaces work, and the role of the default router.

**My evaluation of its helpfulness:**  
Very helpful - clarified the HTTP programming model in Go and helped understand the relationship between handlers, requests, and responses.

### Prompt 3: Go Project Structure
**Prompt used:**  
"How should I organize a Go project and what's the difference between go run and go build?"

**AI's response summary:**  
Explained Go module structure, the purpose of go.mod files, and the difference between running and building Go applications.

**My evaluation of its helpfulness:**  
Helpful - provided context for project organization and build processes, though I needed to experiment with both commands to fully understand the difference.

## 7. Common Issues & Fixes

### Issue 1: "go: cannot find main module"
**Problem:** Trying to run Go files outside a proper module structure  
**Error:** `go: cannot find main module`  
**Fix:** Initialize a Go module first:
```bash
go mod init project-name
```

### Issue 2: Port Already in Use
**Problem:** Another application is using port 8080  
**Error:** `listen tcp :8080: bind: address already in use`  
**Fix:** 
1. Find and stop the process using port 8080
2. Or change the port number to something else like `:8081`

### Issue 3: Go Not Found
**Problem:** Go is not installed or not in PATH  
**Error:** `go: command not found`  
**Fix:** 
1. Install Go from https://golang.org/dl/
2. Add Go to your system PATH
3. Restart your terminal

### Issue 4: Program Starts but Browser Shows Nothing
**Problem:** Server not running or accessing wrong URL  
**Troubleshooting:**
1. Check terminal for startup message
2. Ensure you're visiting http://localhost:8080 (not https)
3. Check if port 8080 is accessible

## 8. References

### Official Documentation
- [Go Official Website](https://golang.org/)
- [Go net/http Package Documentation](https://golang.org/pkg/net/http/)
- [Go Tour Tutorial](https://tour.golang.org/)

### Video Links
- [Go Programming Language Tutorial by TechWithTim](https://www.youtube.com/watch?v=YS4e4q9oBaU)
- [Golang HTTP Server by Web Dev Simplified](https://www.youtube.com/watch?v=yyUHQIec81I)

### Helpful Blog Posts
- [A Simple Go Web Server](https://www.alexedwards.net/blog/a-simple-web-server)
- [Go HTTP Package by Example](https://pkg.go.dev/net/http#example-HandleFunc)
- [Building Web Applications with Go](https://www.mongodb.com/languages/go)

### Stack Overflow References
- [Go HTTP Server Questions](https://stackoverflow.com/questions/tagged/go+http)
- [Common Go HTTP Patterns](https://stackoverflow.com/questions/tagged/go+net-http)

---

## Project Structure

```
go-http-demo/
├── main.go          # Main HTTP server code
├── go.mod           # Go module file
└── README.md        # This documentation
```

## Next Steps

1. **Add More Routes**: Create additional endpoints like `/about`, `/contact`
2. **Add HTML Templates**: Serve dynamic HTML content
3. **Add JSON Responses**: Return JSON instead of plain text
4. **Add Logging**: Implement request logging
5. **Add Middleware**: Add authentication or rate limiting

This toolkit provides a solid foundation for learning Go HTTP programming and can be extended to build more complex web applications.