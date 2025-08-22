// 代码生成时间: 2025-08-23 03:36:23
package main

import (
    "database/sql"
    "fmt"
    "os"
    "log"
    "os/exec"
    "path/filepath"
    "strings"
)

// MigrationBuilder is the struct that holds the necessary information for database migration.
type MigrationBuilder struct {
    MigrationDir string
    DatabaseURL  string
}

// RunMigrations runs all the migration scripts in the specified directory.
func (m *MigrationBuilder) RunMigrations() error {
    // List all files in the migration directory.
    files, err := os.ReadDir(m.MigrationDir)
    if err != nil {
        return err
    }

    // Open the database connection.
    db, err := sql.Open("postgres", m.DatabaseURL)
    if err != nil {
        return err
    }
    defer db.Close()

    // Begin a transaction.
    tx, err := db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()

    for _, file := range files {
        if !file.IsDir() {
            // Construct the full path of the migration file.
            filePath := filepath.Join(m.MigrationDir, file.Name())

            // Execute the migration command.
            cmd := exec.Command("sh", filePath)
            output, err := cmd.CombinedOutput()
            if err != nil {
                log.Printf("Error executing migration %s: %s
", file.Name(), err)
                return err
            }

            // Log the output.
            fmt.Printf("Migration %s executed successfully: %s
", file.Name(), strings.TrimSpace(string(output)))
        }
    }

    // Commit the transaction.
    if err := tx.Commit(); err != nil {
        return err
    }

    return nil
}

func main() {
    // Define the migration directory and database URL.
    migrationDir := "./migrations"
    databaseURL := "postgres://user:password@localhost/dbname"

    // Create a new MigrationBuilder with the provided values.
    migrator := &MigrationBuilder{MigrationDir: migrationDir, DatabaseURL: databaseURL}

    // Run the migrations.
    if err := migrator.RunMigrations(); err != nil {
        log.Fatalf("Failed to run migrations: %s
", err)
    } else {
        fmt.Println("Migrations completed successfully.")
    }
}
