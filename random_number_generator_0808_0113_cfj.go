// 代码生成时间: 2025-08-08 01:13:27
package main

import (
    "fmt"
    "math/rand"
    "time"
    "github.com/spf13/cobra"
)

// RandomNumberGeneratorCmd 表示随机数生成器命令
var RandomNumberGeneratorCmd = &cobra.Command{
    Use:   "random-number-generator",
    Short: "Generate a random number within a specified range",
    Long:  `Generate a random number within a range specified by user.`,
    Args:  cobra.MinimumNArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
        generateRandomNumber(args)
    },
}

// generateRandomNumber 生成指定范围内的随机数
func generateRandomNumber(args []string) {
    // 解析输入参数
    lower, err := strconv.Atoi(args[0])
    if err != nil {
        fmt.Println("Error parsing lower bound: ", err)
        return
    }
    upper, err := strconv.Atoi(args[1])
    if err != nil {
        fmt.Println("Error parsing upper bound: ", err)
        return
    }
    
    // 检查范围有效性
    if lower >= upper {
        fmt.Println("Lower bound must be less than upper bound.")
        return
    }
    
    // 生成随机数
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    randomNumber := r.Intn(upper-lower+1) + lower
    fmt.Println("Random Number: ", randomNumber)
}

func main() {
    // 初始化随机数种子
    rand.Seed(time.Now().UnixNano())
    
    // 添加命令到根命令
    rootCmd := &cobra.Command{
        Use:   "random-number-generator",
        Short: "A brief description of your application",
        Long:  `A longer description that spans multiple lines and likely contains
        Usage examples and options.`,
    }
    rootCmd.AddCommand(RandomNumberGeneratorCmd)
    
    // 执行根命令
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
    }
}