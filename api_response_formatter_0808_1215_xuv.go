// 代码生成时间: 2025-08-08 12:15:09
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "os"

    "github.com/spf13/cobra"
)

// ApiResponse represents the structure of an API response
type ApiResponse struct {
    Success bool        `json:"success"`
    Message string     `json:"message"`
    Data    interface{} `json:"data"`
    Error   string     `json:"error"`
}

// NewApiResponse creates a new ApiResponse instance
func NewApiResponse(success bool, message string, data interface{}, err string) *ApiResponse {
    return &ApiResponse{
        Success: success,
        Message: message,
        Data:    data,
        Error:   err,
    }
}

// PrintJson prints the ApiResponse as a JSON string to stdout
func PrintJson(resp *ApiResponse) error {
    jsonData, err := json.MarshalIndent(resp, "", "  ")
    if err != nil {
        return err
    }
    fmt.Println(string(jsonData))
    return nil
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
    Use:   "api-response-formatter",
    Short: "A tool to format API responses",
    Long:  `A tool to format API responses into JSON format`,
    Run: func(cmd *cobra.Command, args []string) {
        // Example usage: formatting a successful response
        response := NewApiResponse(true, "Operation successful", map[string]string{"key": "value"}, "")
        if err := PrintJson(response); err != nil {
            log.Fatalf("Failed to print JSON: %s", err)
        }

        // Example usage: formatting a response with an error
        responseErr := NewApiResponse(false, "Operation failed", nil, "An error occurred")
        if err := PrintJson(responseErr); err != nil {
            log.Fatalf("Failed to print JSON: %s", err)
        }
    },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func main() {
    Execute()
}
