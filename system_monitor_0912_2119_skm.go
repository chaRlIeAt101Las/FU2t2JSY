// 代码生成时间: 2025-09-12 21:19:39
package main

import (
    "fmt"
    "os"
    "strings"
    "log"
    "golang.org/x/term"
    "github.com/shirou/gopsutil/cpu"
    "github.com/shirou/gopsutil/mem"
    "github.com/shirou/gopsutil/disk"
    "github.com/spf13/cobra"
)

// Define the root command
var rootCmd = &cobra.Command{
    Use:   "system_monitor",
    Short: "A system performance monitoring tool",
    Long:  `A tool to monitor system performance metrics like CPU, memory, and disk usage.`,
    Run: func(cmd *cobra.Command, args []string) {
        monitorSystemPerformance()
    },
}

// Function to monitor system performance
func monitorSystemPerformance() {
    // CPU usage
    cpuPercent, err := cpu.Percent(0, false)
    if err != nil {
        log.Fatalf("Error retrieving CPU usage: %v", err)
    }
    fmt.Printf("CPU Usage: %.2f%%
", cpuPercent[0])

    // Memory usage
    virtualMemory, err := mem.VirtualMemory()
    if err != nil {
        log.Fatalf("Error retrieving memory usage: %v", err)
    }
    fmt.Printf("Memory Usage: %.2f%%
", virtualMemory.UsedPercent)

    // Disk usage
    partitions, err := disk.Partitions(true)
    if err != nil {
        log.Fatalf("Error retrieving disk partitions: %v", err)
    }
    for _, partition := range partitions {
        usage, err := disk.Usage(partition.Mountpoint)
        if err != nil {
            log.Fatalf("Error retrieving disk usage for %s: %v", partition.Mountpoint, err)
        }
        fmt.Printf("Disk Usage for %s: %.2f%%
", partition.Mountpoint, usage.UsedPercent)
    }
}

// Function to initialize and execute the root command
func init() {
    // Add any flags or subcommands to the root command here
    // For example: rootCmd.PersistentFlags().StringVarP(&flagConfig, "config", "c", "default.json", "config file")
}

func main() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
