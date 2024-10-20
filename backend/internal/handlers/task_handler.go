package handlers

import (
    "github.com/gin-gonic/gin"
    "backend/internal/models"
    "backend/internal/service"
    "net/http"
    "strconv"
)

type TaskHandler struct {
    taskService *service.TaskService
}

func NewTaskHandler(taskService *service.TaskService) *TaskHandler {
    return &TaskHandler{taskService: taskService}
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
    var createTask models.TaskCreate
    if err := c.ShouldBindJSON(&createTask); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID := c.GetUint("userID") // From auth middleware
    task, err := h.taskService.CreateTask(&createTask, userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
        return
    }

    var updateTask models.TaskUpdate
    if err := c.ShouldBindJSON(&updateTask); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    task, err := h.taskService.UpdateTask(uint(id), &updateTask)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) GetTasks(c *gin.Context) {
    userID := c.GetUint("userID")
    tasks, err := h.taskService.ListTasks(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
        return
    }
    if err := h.taskService.DeleteTask(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
