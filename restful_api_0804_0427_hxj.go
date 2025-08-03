// 代码生成时间: 2025-08-04 04:27:47
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "log"

    "github.com/spf13/cobra"
)

// Define a simple struct to represent a user.
type User struct {
    Id    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

// APIHandler handles all API requests.
type APIHandler struct {
    users map[int]User
}

// NewAPIHandler creates a new API handler with an empty user map.
func NewAPIHandler() *APIHandler {
    return &APIHandler{users: make(map[int]User)}
}

// GetAllUsers handles GET requests for all users.
func (h *APIHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    err := json.NewEncoder(w).Encode(h.users)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// GetUser handles GET requests for a single user by ID.
func (h *APIHandler) GetUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userID, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    user, ok := h.users[userID]
    if !ok {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    err = json.NewEncoder(w).Encode(user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// AddUser handles POST requests to add a new user.
func (h *APIHandler) AddUser(w http.ResponseWriter, r *http.Request) {
    var user User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    defer r.Body.Close()
    h.users[user.Id] = user
    w.Header().Set("Content-Type", "application/json")
    err = json.NewEncoder(w).Encode(user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func main() {
    handler := NewAPIHandler()

    // Define the routes for the API.
    r := mux.NewRouter()
    r.HandleFunc("/users", handler.GetAllUsers).Methods("GET")
    r.HandleFunc("/users/{id}", handler.GetUser).Methods("GET")
    r.HandleFunc("/users", handler.AddUser).Methods("POST")

    // Define the Cobra command for starting the server.
    var serverCmd = &cobra.Command{
        Use:   "server",
        Short: "Start the RESTful API server",
        Run: func(cmd *cobra.Command, args []string) {
            log.Println("Starting the RESTful API server...")
            if err := http.ListenAndServe(":8080", r); err != nil {
                log.Fatalf("Failed to start server: %v", err)
            }
        },
    }

    // Add the server command to the root command and start the Cobra application.
    rootCmd := &cobra.Command{Use: "api"}
    rootCmd.AddCommand(serverCmd)
    rootCmd.Execute()
}