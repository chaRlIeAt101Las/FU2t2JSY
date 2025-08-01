// 代码生成时间: 2025-08-02 02:33:51
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

// 定义版本号
const version = "1.0.0"

// BackupSyncCmd 定义备份和同步的cobra命令
var BackupSyncCmd = &cobra.Command{
    Use:   "backup-sync [source] [destination]",
    Short: "备份和同步文件工具",
    Long: `一个简单的文件备份和同步工具，用于将指定源目录的文件同步到目标目录。
    可以使用--delete标志来删除目标目录中不存在源目录的文件。`,
    Run:   backupSyncRun,
}

// backupSyncRun 执行备份和同步操作
func backupSyncRun(cmd *cobra.Command, args []string) {
    if len(args) != 2 {
        fmt.Println("源目录和目标目录参数缺失")
        return
    }

    sourceDir := args[0]
    destDir := args[1]

    deleteFlag, _ := cmd.Flags().GetBool("delete")

    // 检查源目录是否存在
    if _, err := os.Stat(sourceDir); os.IsNotExist(err) {
        log.Fatalf("源目录不存在：%s", sourceDir)
    }
    // 检查目标目录是否存在，如果不存在则创建
    if _, err := os.Stat(destDir); os.IsNotExist(err) {
        if err := os.MkdirAll(destDir, 0755); err != nil {
            log.Fatalf("创建目标目录失败：%s", err)
        }
    }

    // 同步文件
    if err := syncFiles(sourceDir, destDir, deleteFlag); err != nil {
        log.Fatalf("同步文件失败：%s", err)
    }
    fmt.Println("文件同步完成")
}

// syncFiles 同步文件到目标目录
func syncFiles(sourceDir, destDir string, deleteFlag bool) error {
    sourceFiles, err := ioutil.ReadDir(sourceDir)
    if err != nil {
        return err
    }

    for _, file := range sourceFiles {
        srcPath := filepath.Join(sourceDir, file.Name())
        destPath := filepath.Join(destDir, file.Name())

        if file.IsDir() {
            if err := syncFiles(srcPath, destPath, deleteFlag); err != nil {
                return err
            }
            continue
        }

        if _, err := os.Stat(destPath); os.IsNotExist(err) {
            if err := copyFile(srcPath, destPath); err != nil {
                return err
            }
        } else if file.ModTime().After(info.ModTime()) {
            if err := copyFile(srcPath, destPath); err != nil {
                return err
            }
        }
    }

    if deleteFlag {
        if err := deleteExtraFiles(destDir, sourceFiles); err != nil {
            return err
        }
    }

    return nil
}

// copyFile 复制文件
func copyFile(srcPath, destPath string) error {
    srcFile, err := os.Open(srcPath)
    if err != nil {
        return err
    }
    defer srcFile.Close()

    destFile, err := os.Create(destPath)
    if err != nil {
        return err
    }
    defer destFile.Close()

    if _, err := io.Copy(destFile, srcFile); err != nil {
        return err
    }
    return nil
}

// deleteExtraFiles 删除目标目录中不存在源目录的文件
func deleteExtraFiles(destDir string, sourceFiles []os.DirEntry) error {
    destFiles, err := ioutil.ReadDir(destDir)
    if err != nil {
        return err
    }

    for _, destFile := range destFiles {
        if destFile.IsDir() {
            continue
        }
        found := false
        for _, srcFile := range sourceFiles {
            if srcFile.Name() == destFile.Name() {
                found = true
                break
            }
        }
        if !found {
            if err := os.Remove(filepath.Join(destDir, destFile.Name())); err != nil {
                return err
            }
        }
    }
    return nil
}

func main() {
    BackupSyncCmd.Flags().BoolP("delete", "d", false, "删除目标目录中不存在源目录的文件")
    BackupSyncCmd.Version = version
    if err := BackupSyncCmd.Execute(); err != nil {
        log.Fatal(err)
    }
}
