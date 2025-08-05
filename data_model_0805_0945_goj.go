// 代码生成时间: 2025-08-05 09:45:45
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "os"
)

// User represents a user data model with fields for UserID, Name, and Email
type User struct {
    UserID  int    `json:"user_id"`
    Name    string `json:"name"`
    Email   string `json:"email"`
    IsAdmin bool   `json:"is_admin"`
}

// CreateUser creates a new user with the provided details
func CreateUser(name string, email string, isAdmin bool) (*User, error) {
    if name == "" || email == "" {
        return nil, fmt.Errorf("name and email are required")
    }

    newUser := User{
        Name: name,
        Email: email,
        IsAdmin: isAdmin,
    }

    // Simulate user ID generation, in real-world scenarios this would be database driven
    newUser.UserID = newUser.generateUserID()

    return &newUser, nil
}

// generateUserID is a placeholder function for generating a unique user ID
func (u *User) generateUserID() int {
    // In a real application, this would involve a more complex logic
    // possibly involving database operations to ensure uniqueness
    return int(u.hashCode())
}

// hashCode is a simple hash function for the email field to simulate user ID generation
func (u *User) hashCode() int {
    // In a real application, use a proper hash function and consider security implications
    return int(hashString(u.Email))
}

// hashString is a simple hash function for demonstration purposes
func hashString(s string) uint32 {
    var hash uint32 = 7
    for i := 0; i < len(s); i++ {
        hash = hash*31 + uint32(s[i])
    }
    return hash
}

// SaveUserToFile saves the user data to a file in JSON format
func SaveUserToFile(user *User, filename string) error {
    file, err := os.Create(filename)
    if err != nil {
        return fmt.Errorf("error creating file: %w", err)
    }
    defer file.Close()

    data, err := json.MarshalIndent(user, "", "    ")
    if err != nil {
        return fmt.Errorf("error marshalling user data: %w", err)
    }

    _, err = file.Write(data)
    if err != nil {
        return fmt.Errorf("error writing to file: %w", err)
    }

    return nil
}

// main function to demonstrate creating and saving a user
func main() {
    user, err := CreateUser("John Doe", "john.doe@example.com", false)
    if err != nil {
        log.Fatalf("Error creating user: %s", err)
    }

    fmt.Printf("Created user: %+v
", user)

    err = SaveUserToFile(user, "user.json")
    if err != nil {
        log.Fatalf("Error saving user to file: %s", err)
    }

    fmt.Println("User saved successfully")
}
