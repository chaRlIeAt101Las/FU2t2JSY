// 代码生成时间: 2025-08-29 03:10:05
package main

import (
    "fmt"
    "time"
    "github.com/robfig/cron/v3"
    "log"
)

// Scheduler is a struct that holds the cron scheduler
type Scheduler struct {
    cron *cron.Cron
}

// NewScheduler creates a new Scheduler instance
func NewScheduler() *Scheduler {
    return &Scheduler{
        cron: cron.New(cron.WithSeconds()), // 创建一个具有秒级别的cron调度器
    }
}

// AddTask adds a new task to the scheduler
func (s *Scheduler) AddTask(schedule string, task func()) error {
    _, err := s.cron.AddFunc(schedule, task)
    if err != nil {
        return err
    }
    return nil
}

// Start starts the scheduler
func (s *Scheduler) Start() {
    s.cron.Start()
}

// Stop stops the scheduler
func (s *Scheduler) Stop() {
    s.cron.Stop()
}

// TaskExample is a sample task function that prints the current time every minute
func TaskExample() {
    fmt.Println("Task executed at", time.Now().Format("2006-01-02 15:04:05"))
}

func main() {
    scheduler := NewScheduler()
    defer scheduler.Stop()

    // Add a task that will run every minute
    if err := scheduler.AddTask("* * * * *", TaskExample); err != nil {
        log.Fatalf("Failed to add task: %v", err)
    }

    // Wait forever to allow tasks to run
    select {}
}
