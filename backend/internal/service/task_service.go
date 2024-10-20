package service

import (
    "backend/internal/models"
    "backend/internal/repository"
)

type TaskService struct {
    repo *repository.Repository
}

func NewTaskService(repo *repository.Repository) *TaskService {
    return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(create *models.TaskCreate, userID uint) (*models.Task, error) {
    task := &models.Task{
        Title:       create.Title,
        Description: create.Description,
        DueDate:     create.DueDate,
        Priority:    create.Priority,
        Status:      "pending",
        UserID:      userID,
    }

    err := s.repo.CreateTask(task)
    if err != nil {
        return nil, err
    }

    return task, nil
}

func (s *TaskService) UpdateTask(id uint, update *models.TaskUpdate) (*models.Task, error) {
    task, err := s.repo.GetTaskByID(id)
    if err != nil {
        return nil, err
    }

    if update.Title != nil {
        task.Title = *update.Title
    }
    if update.Description != nil {
        task.Description = *update.Description
    }
    if update.DueDate != nil {
        task.DueDate = *update.DueDate
    }
    if update.Priority != nil {
        task.Priority = *update.Priority
    }
    if update.Status != nil {
        task.Status = *update.Status
    }

    err = s.repo.UpdateTask(task)
    if err != nil {
        return nil, err
    }

    return task, nil
}


func (s *TaskService) SearchTasks(userID uint, query *models.TaskSearchQuery) ([]models.Task, error) {
    db := s.repo.DB().Where("user_id = ?", userID)

    if query.Status != "" {
        db = db.Where("status = ?", query.Status)
    }

    if query.Priority != "" {
        db = db.Where("priority = ?", query.Priority)
    }

    if !query.StartDate.IsZero() {
        db = db.Where("due_date >= ?", query.StartDate)
    }

    if !query.EndDate.IsZero() {
        db = db.Where("due_date <= ?", query.EndDate)
    }

    if query.SearchTerm != "" {
        searchTerm := "%" + query.SearchTerm + "%"
        db = db.Where("title LIKE ? OR description LIKE ?", searchTerm, searchTerm)
    }

    var tasks []models.Task
    err := db.Find(&tasks).Error
    return tasks, err
}

func (s *TaskService) ListTasks(userID uint) ([]models.Task, error) {
    tasks, err := s.repo.ListTasks(userID)
    if err != nil {
        return nil, err
    }
    return tasks, nil
}

func (s *TaskService) DeleteTask(taskID uint) error {
    return s.repo.DeleteTask(taskID)
}

