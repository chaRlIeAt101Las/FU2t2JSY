// 代码生成时间: 2025-08-05 15:56:17
package main

import (
    "fmt"
    "time"
    "github.com/robfig/cron/v3"
)

// Scheduler 定义定时任务调度器
type Scheduler struct {
    cron *cron.Cron
}

// NewScheduler 创建一个新的定时任务调度器实例
func NewScheduler() *Scheduler {
    return &Scheduler{
        cron: cron.New(cron.WithSeconds()),
    }
}

// AddJob 添加一个定时任务
func (s *Scheduler) AddJob(spec string, cmd func()) error {
    if _, err := s.cron.AddFunc(spec, cmd); err != nil {
        return fmt.Errorf("failed to add job: %w", err)
    }
    return nil
}

// Start 开始运行定时任务调度器
func (s *Scheduler) Start() {
    s.cron.Start()
}

// Stop 停止运行定时任务调度器
func (s *Scheduler) Stop() error {
    return s.cron.Stop()
}

// ExampleJob 一个示例定时任务函数
func ExampleJob() {
    fmt.Println("Example job is running...")
}

func main() {
    // 创建调度器实例
    scheduler := NewScheduler()

    // 添加示例定时任务，每5秒执行一次
    if err := scheduler.AddJob("*/5 * * * * * *", ExampleJob); err != nil {
        fmt.Printf("Error adding job: %s
", err)
        return
    }

    fmt.Println("Scheduler is running...
Press CTRL+C to stop")

    // 启动调度器
    scheduler.Start()

    // 阻塞主线程，等待调度器停止
    select {}
}
