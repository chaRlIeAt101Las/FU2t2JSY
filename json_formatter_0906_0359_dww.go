// 代码生成时间: 2025-09-06 03:59:00
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "os"

    "github.com/spf13/cobra"
)

// JSONFormatterCmd represents the base command for the JSON formatter when called without any subcommands
var JSONFormatterCmd = &cobra.Command{
    Use:   "json-formatter",
    Short: "A brief description of your application",
    Long: `A longer description that spans multiple lines and
    likely contains examples to show how to use the application.`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("JSON Formatter called")
        inputJSON, err := cmd.Flags().GetString("input")
        if err != nil {
            log.Fatalf("Error reading input flag: %s", err)
        }

        var jsonData interface{}
        if err := json.Unmarshal([]byte(inputJSON), &jsonData); err != nil {
            log.Fatalf("Error parsing JSON: %s", err)
        }

        formattedJSON, err := json.MarshalIndent(jsonData, "", "  ")
        if err != nil {
            log.Fatalf("Error formatting JSON: %s", err)
        }

        fmt.Println(string(formattedJSON))
    },
}

// init registers the JSONFormatterCmd with the provided cobra command
func init() {
    JSONFormatterCmd.Flags().StringP("input", "i", "", "Input JSON to be formatted")
}

func main() {
    if err := JSONFormatterCmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, "Error executing the command: ", err)
        os.Exit(1)
    }
}