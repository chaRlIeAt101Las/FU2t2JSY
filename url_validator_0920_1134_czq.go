// 代码生成时间: 2025-09-20 11:34:59
package main

import (
    "fmt"
    "net/url"
    "strings"
    "os"
    "github.com/spf13/cobra"
)

// URLValidatorCmd is a Cobra command to validate the URL
var URLValidatorCmd = &cobra.Command{
    Use:   "url-validator <URL>",
    Short: "Validates the given URL",
    Long: `A utility to check if a URL is valid
   Simply provide the URL as an argument to the command.`,
    Args: cobra.ExactArgs(1), // This command needs exactly one argument
    RunE: validateURL, // Function to execute when the command is run
}

// validateURL is the function that performs the URL validation
func validateURL(cmd *cobra.Command, args []string) error {
    // Parse the URL argument
    inputURL := args[0]
    u, err := url.ParseRequestURI(inputURL)
    if err != nil {
        return err // Return the error if parsing fails
    }

    // Check for valid scheme and host
    if len(u.Scheme) == 0 || len(u.Host) == 0 {
        return fmt.Errorf("invalid URL: missing scheme or host")
    }

    // Print the result of the validation
    fmt.Printf("The URL '%s' is valid.
", inputURL)
    return nil // Return nil if the URL is valid
}

// main is the entry point of the application
func main() {
    // Add the URLValidatorCmd to the root command
    rootCmd := &cobra.Command{Use: "url-validator"}
    rootCmd.AddCommand(URLValidatorCmd)

    // Execute the root command, which will run the URLValidatorCmd if provided
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err) // Print the error if the command execution fails
        os.Exit(1) // Exit with a non-zero status code
    }
}
