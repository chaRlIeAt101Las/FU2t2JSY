// 代码生成时间: 2025-09-12 11:47:03
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "embed"
    "golang.org/x/exp/slices"
)

//go:embed templates/*
var templates embed.FS

//go:embed static/*
var staticFiles embed.FS

// Layout defines the structure for a page layout
type Layout struct {
    BasePath string
    Assets   http.FileSystem
}

// NewLayout creates a new instance of Layout
func NewLayout(basePath string) *Layout {
    return &Layout{
        BasePath: basePath,
        Assets:   staticFiles,
    }
}

// ServeHTTP serves the HTTP request for the layout
func (l *Layout) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // Handle GET requests only
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Redirect to the base path if path is empty
    if r.URL.Path == l.BasePath {
        http.Redirect(w, r, l.BasePath+"index.html", http.StatusFound)
        return
    }

    // Serve the requested file from the static assets
    filepath := r.URL.Path[len(l.BasePath):]
    if filepath == "" {
        filepath = "index.html"
    }
    http.StripPrefix(l.BasePath, http.FileServer(http.FS(l.Assets))).ServeHTTP(w, r)
}

func main() {
    // Set up the layout with the base path
    layout := NewLayout("/")

    // Set up the HTTP server
    http.Handle(layout.BasePath, http.StripPrefix(layout.BasePath, layout))

    // Set up the template server
    http.Handle("/templates/", http.FileServer(http.FS(templates)))

    // Set up the port and start the server
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    log.Printf("Server is running on http://localhost:%s", port)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
