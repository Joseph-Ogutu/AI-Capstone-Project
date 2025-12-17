package main

import (
    "fmt"
    "log"
    "net/http"
)

// homeHandler handles requests to the home page
func homeHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    
    w.Header().Set("Content-Type", "text/html")
    fmt.Fprintf(w, `
    <html>
        <head><title>Go HTTP Server</title></head>
        <body>
            <h1>Hello, World from Go!</h1>
            <p>Welcome to your first Go HTTP server.</p>
            <p><a href="/about">About</a> | <a href="/api/time">Current Time</a></p>
        </body>
    </html>
    `)
}

// aboutHandler provides information about this Go server
func aboutHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    fmt.Fprintf(w, `
    <html>
        <head><title>About - Go HTTP Server</title></head>
        <body>
            <h1>About This Server</h1>
            <p>This is a simple HTTP server built with Go's net/http package.</p>
            <p>Features:</p>
            <ul>
                <li>Basic routing</li>
                <li>HTML responses</li>
                <li>JSON API endpoints</li>
            </ul>
            <p><a href="/">Back to Home</a></p>
        </body>
    </html>
    `)
}

// timeHandler returns current time as JSON
func timeHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, `{"message": "Current time from Go server", "status": "success"}`)
}

func main() {
    // Register HTTP handlers
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/about", aboutHandler)
    http.HandleFunc("/api/time", timeHandler)
    
    // Start the server
    serverAddress := ":8080"
    log.Printf("Server starting on http://localhost%s", serverAddress)
    
    // Start HTTP server - ListenAndServe always returns an error
    err := http.ListenAndServe(serverAddress, nil)
    if err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}