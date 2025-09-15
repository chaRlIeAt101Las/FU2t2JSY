// 代码生成时间: 2025-09-15 12:20:53
package main

import (
    "encoding/json"
    "fmt"
    "github.com/spf13/cobra"
    "os"
)

// User represents a user data model
type User struct {
    ID      int    `json:"id"`
    Name    string `json:"name"`
    Email   string `json:"email"`
    Age     int    `json:"age"`
}

// NewUser creates a new user instance
func NewUser(id int, name, email string, age int) *User {
    return &User{
        ID:   id,
        Name: name,
        Email: email,
        Age:  age,
    }
}

// Validate checks if the user data is valid
func (u *User) Validate() error {
    if u.Name == "" || u.Email == "" || u.Age <= 0 {
        return fmt.Errorf("user data is invalid")
    }
    return nil
}

// toJSON converts user data to JSON format
func (u *User) toJSON() ([]byte, error) {
    return json.Marshal(u)
}

// CreateUserCommand is a Cobra command to create a user
var CreateUserCommand = &cobra.Command{
    Use:   "create",
    Short: "Create a new user",
    Long:  `Create a new user with provided data`,
    Run: func(cmd *cobra.Command, args []string) {
        if len(args) < 4 {
            fmt.Println("Error: Please provide id, name, email, and age")
            return
        }

        id, err := strconv.Atoi(args[0])
        if err != nil {
            fmt.Printf("Error converting id to int: %s
", err)
            return
        }

        name := args[1]
        email := args[2]
        age, err := strconv.Atoi(args[3])
        if err != nil {
            fmt.Printf("Error converting age to int: %s
", err)
            return
        }

        user := NewUser(id, name, email, age)
        if err := user.Validate(); err != nil {
            fmt.Printf("Error validating user: %s
", err)
            return
        }

        jsonBytes, err := user.toJSON()
        if err != nil {
            fmt.Printf("Error converting user to JSON: %s
", err)
            return
        }

        fmt.Printf("User created: %s
", string(jsonBytes))
    },
}

func main() {
    CreateUserCommand.Flags().IntP("id", "i", 0, "User ID")
    CreateUserCommand.Flags().StringP("name", "n", "", "User Name")
    CreateUserCommand.Flags().StringP("email", "e", "", "User Email")
    CreateUserCommand.Flags().IntP("age", "a", 0, "User Age")

    if err := CreateUserCommand.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, "Error executing command: ", err)
        os.Exit(1)
    }
}
