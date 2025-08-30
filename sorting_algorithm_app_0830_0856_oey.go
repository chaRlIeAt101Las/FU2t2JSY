// 代码生成时间: 2025-08-30 08:56:42
package main

import (
    "fmt"
    "os"
    "sort"

    "github.com/spf13/cobra"
)

// Version number
const Version = "1.0.0"

// SortingAlgorithmCmd represents the sorting command
var SortingAlgorithmCmd = &cobra.Command{
    Use:   "sort",
    Short: "Sorts an array of integers",
    Long:  `This command sorts an array of integers using the specified algorithm.`,
    Args:  cobra.MinimumNArgs(1), // Expects at least one argument
    Run:   sortNumbers,
}

// sortNumbers is the implementation of the 'sort' command
func sortNumbers(cmd *cobra.Command, args []string) {
    // Convert string arguments to integers
    numbers, err := convertStringsToInts(args)
    if err != nil {
        fmt.Printf("Error converting arguments to integers: %s
", err)
        os.Exit(1)
    }

    // Sort the numbers using the standard library
    sort.Ints(numbers)

    // Print the sorted numbers
    for _, num := range numbers {
        fmt.Printf("%d ", num)
    }
    fmt.Println()
}

// convertStringsToInts converts a slice of strings to a slice of integers
func convertStringsToInts(args []string) ([]int, error) {
    var numbers []int
    for _, strNum := range args {
        num, err := strconv.Atoi(strNum)
        if err != nil {
            return nil, fmt.Errorf("invalid integer: %s", strNum)
        }
        numbers = append(numbers, num)
    }
    return numbers, nil
}

func main() {
    // Create a new cobra command with the version
    rootCmd := &cobra.Command{
        Use:   "sorting-app",
        Short: "A simple sorting algorithm application",
        Version: Version,
    }

    // Add the sorting command to the root command
    rootCmd.AddCommand(SortingAlgorithmCmd)

    // Execute the root command
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
