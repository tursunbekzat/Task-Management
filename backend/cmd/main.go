package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"backend/internal/handlers"
	"backend/internal/middleware"
	"backend/internal/repository"
	"backend/internal/service"

	"net/http"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Build DSN from environment variables
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET not set in .env")
	}

	// Initialize repositories and services
	repo, err := repository.NewRepository(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	userRepo := repository.NewUserRepository(repo.DB())

	userService := service.NewUserService(userRepo, jwtSecret)
	taskService := service.NewTaskService(repo)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userService)
	taskHandler := handlers.NewTaskHandler(taskService)

	// Set up the Gin router
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, X-Requested-With, Accept")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})

	router.Use(middleware.LoggingMiddleware())

	// Define routes
	api := router.Group("/api")
	{
		api.POST("/auth/register", userHandler.Register)
		api.POST("/auth/login", userHandler.Login)

		authGroup := api.Group("/tasks")
		authGroup.Use(middleware.AuthMiddleware(jwtSecret))
		{
			authGroup.POST("", taskHandler.CreateTask)
			authGroup.GET("", taskHandler.GetTasks)
			authGroup.PUT("/:id", taskHandler.UpdateTask)
			authGroup.DELETE("/:id", taskHandler.DeleteTask)
		}
	}

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on port %s...", port)
	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
