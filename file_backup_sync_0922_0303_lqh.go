// 代码生成时间: 2025-09-22 03:03:24
package main

import (
    "bufio"
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "time"

    "github.com/spf13/cobra"
)

// FileInfo holds the file information such as name, size, and hash.
type FileInfo struct {
    Name    string
    Size    int64
    Hash    string
    ModTime time.Time
}

// backupFile copies the source file to the destination.
func backupFile(src, dst string) error {
    srcFile, err := os.Open(src)
    if err != nil {
        return err
    }
    defer srcFile.Close()

    dstFile, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer dstFile.Close()

    _, err = io.Copy(dstFile, srcFile)
    return err
}

// calculateFileHash calculates the SHA256 hash of a file.
func calculateFileHash(filePath string) (string, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return "", err
    }
    defer file.Close()

    hasher := sha256.New()
    if _, err := io.Copy(hasher, file); err != nil {
        return "", err
    }
    return hex.EncodeToString(hasher.Sum(nil)), nil
}

// syncFiles synchronizes files between source and destination.
func syncFiles(src, dst string) error {
    srcFiles, err := ioutil.ReadDir(src)
    if err != nil {
        return err
    }

    for _, file := range srcFiles {
        srcPath := filepath.Join(src, file.Name())
        dstPath := filepath.Join(dst, file.Name())

        srcInfo, err := os.Stat(srcPath)
        if err != nil {
            return err
        }

        if _, err := os.Stat(dstPath); os.IsNotExist(err) {
            // File does not exist in destination, copy it.
            if err := backupFile(srcPath, dstPath); err != nil {
                return err
            }
        } else if srcInfo.ModTime().After(file.ModTime()) {
            // Source file is newer, overwrite it.
            if err := backupFile(srcPath, dstPath); err != nil {
                return err
            }
        }
    }
    return nil
}

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
    Use:   "file_backup_sync",
    Short: "A brief description of your application",
    Long: `A longer description that spans multiple lines
and likely contains examples to show how to use the application.`,
    // Uncomment the following line if your bare application
    // has an action associated with it:
    // Run: func(cmd *cobra.Command, args []string) {
    // },
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
    // Add your commands here.
    //cobra.OnInitialize(initConfig)

    // Create backup command.
    backupCmd := &cobra.Command{
        Use:   "backup [source] [destination]",
        Short: "Backup files from source to destination",
        Long:  "Copy files from the source directory to the destination directory",
        Run: func(cmd *cobra.Command, args []string) {
            if len(args) != 2 {
                fmt.Println("Error: Two arguments are required.")
                return
            }
            src, dst := args[0], args[1]
            if err := syncFiles(src, dst); err != nil {
                fmt.Printf("Error syncing files: %v
", err)
            } else {
                fmt.Println("Backup completed successfully.")
            }
        },
    }
    rootCmd.AddCommand(backupCmd)

    if err := rootCmd.Execute(); err != nil {
        log.Fatalf("Error executing the command: %v
", err)
    }
}
