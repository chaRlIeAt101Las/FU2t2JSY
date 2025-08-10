// 代码生成时间: 2025-08-11 04:00:08
package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"
    "time"

    "github.com/spf13/cobra"
)

// 定义 FolderOrganizer 结构体用于封装文件夹整理逻辑
type FolderOrganizer struct {
    RootPath string
}

// NewFolderOrganizer 创建 FolderOrganizer 实例
func NewFolderOrganizer(rootPath string) *FolderOrganizer {
    return &FolderOrganizer{
        RootPath: rootPath,
    }
}

// Organize 执行文件夹整理工作
func (f *FolderOrganizer) Organize() error {
    // 读取 rootPath 下的所有文件和文件夹
    files, err := os.ReadDir(f.RootPath)
    if err != nil {
        return fmt.Errorf("failed to read directory: %w", err)
    }

    for _, file := range files {
        filePath := filepath.Join(f.RootPath, file.Name())
        if file.IsDir() {
            // 如果是文件夹，则递归整理
            err = f.Organize()
            if err != nil {
                return fmt.Errorf("failed to organize directory %s: %w", filePath, err)
            }
        } else {
            // 如果是文件，根据文件名进行分类并移动到对应的文件夹
            err = f.categorizeFile(filePath)
            if err != nil {
                return fmt.Errorf("failed to categorize file %s: %w", filePath, err)
            }
        }
    }
    return nil
}

// categorizeFile 根据文件名分类文件并移动到对应的文件夹
func (f *FolderOrganizer) categorizeFile(filePath string) error {
    fileName := filepath.Base(filePath)
    // 确定文件分类
    category := "others"
    if strings.HasSuffix(fileName, ".jpg") || strings.HasSuffix(fileName, ".png\) {
        category = "images"
    } else if strings.HasSuffix(fileName, ".mp3") || strings.HasSuffix(fileName, ".wav") {
        category = "audio"
    }
    // ... 可以根据需要添加更多分类

    // 创建分类文件夹
    categoryPath := filepath.Join(f.RootPath, category)
    if _, err := os.Stat(categoryPath); os.IsNotExist(err) {
        err = os.Mkdir(categoryPath, 0755)
        if err != nil {
            return fmt.Errorf("failed to create directory %s: %w", categoryPath, err)
        }
    }

    // 移动文件到对应的分类文件夹
    destPath := filepath.Join(categoryPath, fileName)
    if err := os.Rename(filePath, destPath); err != nil {
        return fmt.Errorf("failed to move file %s to %s: %w", filePath, destPath, err)
    }
    fmt.Printf("Moved file %s to %s\
", filePath, destPath)
    return nil
}

// main 函数执行程序的主逻辑
func main() {
    var rootPath string
    var useLongDateFormat bool

    // 创建 cobra 命令
    rootCmd := &cobra.Command{
        Use:   "folder-organizer",
        Short: "Organize files in a given directory",
        Run: func(cmd *cobra.Command, args []string) {
            // 检查是否提供了文件夹路径
            if rootPath == "" {
                fmt.Println("Error: No root path provided")
                return
            }

            // 创建文件夹整理器并执行整理
            organizer := NewFolderOrganizer(rootPath)
            if err := organizer.Organize(); err != nil {
                log.Fatalf("Error organizing folder: %v", err)
            }
        },
    }

    // 添加命令行参数
    rootCmd.PersistentFlags().StringVarP(&rootPath, "root-path", "r", ".", "Root directory path for organizing files")
    rootCmd.PersistentFlags().BoolVarP(&useLongDateFormat, "long-date", "l", false, "Use long date format for log entries")

    // 执行命令
    if err := rootCmd.Execute(); err != nil {
        log.Fatalf("Error executing command: %v", err)
    }
}
