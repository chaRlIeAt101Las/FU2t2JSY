// 代码生成时间: 2025-09-11 01:37:31
package main

import (
    "crypto/sha256"
    "fmt"
    "io/ioutil"
    "os"
    "strings"
    "golang.org/x/crypto/sha3"
    "github.com/spf13/cobra"
)

// Command represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
    Use:   "hash_calculator",
    Short: "A brief description of your application",
    Long: `A longer description that spans multiple lines and likely contains
examples of how to use the application.`,
    Run: func(cmd *cobra.Command, args []string) {
        calculateHash()
    },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main().
// It only needs to happen once to the rootCmd.
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func main() {
    Execute()
}

// calculateHash function calculates the hash for the provided input
func calculateHash() {
    var input string
    fmt.Print("Please enter a message to hash: ")
    fmt.Scanln(&input)

    // Calculate SHA256 hash
    sha256Hash := sha256.Sum256([]byte(input))
    fmt.Printf("SHA256 Hash: %x
", sha256Hash)

    // Calculate SHA3-256 hash
    sha3Hash := sha3.Sum256([]byte(input))
    fmt.Printf("SHA3-256 Hash: %x
", sha3Hash)
}
