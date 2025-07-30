// 代码生成时间: 2025-07-31 00:38:23
// unittest_framework.go

package main

import (
    "fmt"
    "testing"
    "os"
    "path/filepath"
    "strings"

    // 引入Cobra库
    "github.com/spf13/cobra"
)

// 定义测试命令
var testCmd = &cobra.Command{
    Use:   "test <testfile>",
    Short: "Run unit tests",
    Long:  "Run unit tests for a given test file or all test files if none specified",
    Args:  cobra.MaximumNArgs(1),
    RunE:  runTests,
}

// 添加测试命令到rootCmd
func init() {
    rootCmd.AddCommand(testCmd)
}

// runTests 函数执行测试
func runTests(cmd *cobra.Command, args []string) error {
    // 检查参数
    if len(args) == 0 {
        // 测试所有文件
        return runAllTests()
    } else {
        // 测试单个文件
        return runTestFile(args[0])
    }
}

// runAllTests 运行所有测试文件
func runAllTests() error {
    fmt.Println("Running all tests...")
    // 遍历测试文件夹
    err := filepath.WalkDir("./tests", visitFunc)
    if err != nil {
        return err
    }
    return nil
}

// runTestFile 运行单个测试文件
func runTestFile(filename string) error {
    fmt.Printf("Running tests in %s...
", filename)
    // 运行测试
    err := testing.Main(strings.TrimSuffix(filename, "_test.go"))
    if err != nil {
        return err
    }
    return nil
}

// visitFunc 是filepath.WalkDir的回调函数，用于遍历文件
func visitFunc(path string, d os.DirEntry, err error) error {
    if err != nil {
        return err
    }
    if d.IsDir() {
        return nil
    }
    if strings.HasSuffix(path, "_test.go") {
        // 运行单个测试文件
        err = runTestFile(path)
        if err != nil {
            return err
        }
    }
    return nil
}

// main 函数设置Cobra并执行命令
func main() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}

// rootCmd 是应用的根命令
var rootCmd = &cobra.Command{
    Use:   "unittest",
    Short: "Unit test framework",
    Long:  `Run unit tests for your Go applications`,
}
