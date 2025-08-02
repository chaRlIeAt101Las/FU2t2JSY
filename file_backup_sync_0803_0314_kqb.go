// 代码生成时间: 2025-08-03 03:14:30
package main

import (
    "bufio"
    "flag"
    "fmt"
    "io"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"
    "time"
)

// 定义一个结构体，用于存储文件备份和同步的配置信息
type Config struct {
    Source string
    Destination string
    Exclude []string
}

// 检查文件是否应该被排除
func shouldExclude(filePath string, excludePatterns []string) bool {
    for _, pattern := range excludePatterns {
        if matched, _ := filepath.Match(pattern, filepath.Base(filePath)); matched {
            return true
        }
    }
    return false
}

// 递归地遍历源文件夹，同步文件到目标文件夹
func syncFiles(source, destination string, excludePatterns []string) error {
    err := filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if info.IsDir() && shouldExclude(path, excludePatterns) {
            return filepath.SkipDir
        }
        if !info.IsDir() && shouldExclude(path, excludePatterns) {
            return nil
        }
        relPath, err := filepath.Rel(source, path)
        if err != nil {
            return err
        }
        destPath := filepath.Join(destination, relPath)
        if info.IsDir() {
            _, err = os.Stat(destPath)
            if os.IsNotExist(err) {
                err = os.MkdirAll(destPath, os.ModePerm)
                if err != nil {
                    return err
                }
            }
        } else {
            err = copyFile(path, destPath)
            if err != nil {
                return err
            }
        }
        return nil
    })
    return err
}

// 复制单个文件
func copyFile(src, dst string) error {
    sourceFile, err := os.Open(src)
    if err != nil {
        return err
    }
    defer sourceFile.Close()

    destinationFile, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer destinationFile.Close()

    _, err = io.Copy(destinationFile, sourceFile)
    return err
}

func main() {
    var config Config
    flag.StringVar(&config.Source, "source", "", "源文件夹路径")
    flag.StringVar(&config.Destination, "destination", "", "目标文件夹路径")
    flag.StringSliceVar(&config.Exclude, "exclude", []string{}, "排除的文件或文件夹模式")
    flag.Parse()

    if config.Source == "" || config.Destination == "" {
        fmt.Println("源文件夹和目标文件夹路径不能为空")
        return
    }

    err := syncFiles(config.Source, config.Destination, config.Exclude)
    if err != nil {
        fmt.Printf("同步文件时发生错误: %v", err)
    } else {
        fmt.Println("文件同步完成")
    }
}
