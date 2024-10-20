package models

import (
    "time"
    "gorm.io/gorm"
)

type Priority string

const (
    PriorityLow    Priority = "low"
    PriorityMedium Priority = "medium"
    PriorityHigh   Priority = "high"
)

type Task struct {
    gorm.Model
    ID          uint      `json:"ID" gorm:"primaryKey"`
    Title       string    `json:"title" gorm:"not null"`
    Description string    `json:"description"`
    DueDate     time.Time `json:"dueDate"`
    Priority    Priority  `json:"priority"`
    Status      string    `json:"status"`
    UserID      uint      `json:"userId"`
    User        User      `json:"user"`
}

type TaskCreate struct {
    Title       string    `json:"title" binding:"required"`
    Description string    `json:"description"`
    DueDate     time.Time `json:"dueDate" binding:"required"`
    Priority    Priority  `json:"priority" binding:"required"`
}

type TaskUpdate struct {
    Title       *string    `json:"title"`
    Description *string    `json:"description"`
    DueDate     *time.Time `json:"dueDate"`
    Priority    *Priority  `json:"priority"`
    Status      *string    `json:"status"`
}

type TaskSearchQuery struct {
    Status     string
    Priority   string
    StartDate  time.Time
    EndDate    time.Time
    SearchTerm string
}