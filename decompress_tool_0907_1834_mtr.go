// 代码生成时间: 2025-09-07 18:34:46
package main

import (
    "archive/zip"
    "flag"
    "fmt"
    "io"
    "os"
    "path/filepath"
)

// Decompress decompresses the archive file into the specified directory.
func Decompress(src, dest string) error {
    reader, err := zip.OpenReader(src)
    if err != nil {
        return err
    }
    defer reader.Close()

    for _, file := range reader.File {
        filePath := filepath.Join(dest, file.Name)
        if file.FileInfo().IsDir() {
            // Create directory if it does not exist
            if err = os.MkdirAll(filePath, os.ModePerm); err != nil {
                return err
            }
        } else {
            // Create directory if it does not exist
            if err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
                return err
            }
            outFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
            if err != nil {
                return err
            }
            defer outFile.Close()

            fileInArchive, err := file.Open()
            if err != nil {
                return err
            }
            defer fileInArchive.Close()
            
            _, err = io.Copy(outFile, fileInArchive)
            if err != nil {
                return err
            }
        }
    }
    return nil
}

func main() {
    // Define flags for the source and destination paths.
    src := flag.String("src", "", "The source zip file path")
    dest := flag.String("dest", "", "The destination directory path")
    flag.Parse()

    if *src == "" || *dest == "" {
        fmt.Println("Both source and destination paths are required.")
        return
    }

    // Decompress the file into the specified directory.
    if err := Decompress(*src, *dest); err != nil {
        fmt.Printf("An error occurred: %v", err)
    } else {
        fmt.Println("Decompression successful.")
    }
}
