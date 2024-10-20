package handlers

import (
    "github.com/gin-gonic/gin"
    "backend/internal/models"
    "backend/internal/service"
    "net/http"
)

type UserHandler struct {
    userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
    return &UserHandler{userService: userService}
}

func (h *UserHandler) Register(c *gin.Context) {
    var createUser models.UserCreate
    if err := c.ShouldBindJSON(&createUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := h.userService.Register(&createUser)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) Login(c *gin.Context) {
    var login models.UserLogin
    if err := c.ShouldBindJSON(&login); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    token, err := h.userService.Login(&login)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
    userID := c.GetUint("userID") // From auth middleware
    
    var update models.UserUpdate
    if err := c.ShouldBindJSON(&update); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := h.userService.UpdateUser(userID, &update)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, user)
}