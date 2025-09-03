// 代码生成时间: 2025-09-03 21:08:17
package main

import (
    "fmt"
    "io"
    "log"
    "os"
    "path/filepath"
    "time"

    "github.com/spf13/cobra"
)

// 文件备份和同步工具
var rootCmd = &cobra.Command{
    Use:   "fileBackupSync", // 命令名称
    Short: "A tool for file backup and synchronization", // 命令简短描述
    Long: `
A tool for file backup and synchronization. It allows you to backup
and synchronize files between different directories.`, // 命令详细描述
    Run: func(cmd *cobra.Command, args []string) {
        // 执行备份和同步操作
        sourcePath, _ := cmd.Flags().GetString("source")
        destPath, _ := cmd.Flags().GetString("destination\)

        if sourcePath == "" || destPath == "" {
            fmt.Println("Source and destination paths must be provided.")
            return
        }

        err := syncFiles(sourcePath, destPath)
        if err != nil {
            log.Fatalf("Error syncing files: %v", err)
        }
    },
}

// syncFiles 同步文件从 sourcePath 到 destPath
func syncFiles(sourcePath, destPath string) error {
    // 获取 sourcePath 下的所有文件
    files, err := os.ReadDir(sourcePath)
    if err != nil {
        return fmt.Errorf("failed to read source directory: %w", err)
    }

    for _, file := range files {
        srcFile := filepath.Join(sourcePath, file.Name())
        destFile := filepath.Join(destPath, file.Name())

        // 如果目标文件不存在，则创建它
        if _, err := os.Stat(destFile); os.IsNotExist(err) {
            err := os.WriteFile(destFile, []byte{}, 0644)
            if err != nil {
                return fmt.Errorf("failed to create file at destination: %w", err)
            }
        }

        // 读取源文件内容并写入目标文件
        src, err := os.ReadFile(srcFile)
        if err != nil {
            return fmt.Errorf("failed to read source file: %w", err)
        }

        err = os.WriteFile(destFile, src, 0644)
        if err != nil {
            return fmt.Errorf("failed to write to destination file: %w", err)
        }

        // 更新最后修改时间
        err = os.Chtimes(destFile, time.Now(), time.Now())
        if err != nil {
            return fmt.Errorf("failed to update modification time: %w", err)
        }
    }

    return nil
}

func main() {
    // 添加命令行参数
    rootCmd.Flags().StringP("source", "s", "", "Source directory for file backup and synchronization")
    rootCmd.Flags().StringP("destination\, "d", "", "Destination directory for file backup and synchronization")

    // 执行根命令
    if err := rootCmd.Execute(); err != nil {
        log.Fatalf("Error executing command: %v", err)
    }
}
