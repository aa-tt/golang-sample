package main

import (
	"log"
	"microservice/handlers"
	"microservice/middleware"
	"microservice/models"
	"microservice/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Database (Mocking H2 with SQLite in-memory)
	if err := models.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Start Background Consumer
	service.StartConsumer()

	// Initialize Gin Router
	r := gin.New()

	// Global Middlewares
	r.Use(gin.Recovery())
	r.Use(middleware.LoggingMiddleware())
	r.Use(handlers.ErrorHandler())

	// Public Routes
	r.POST("/login", handlers.Login)

	// API Gateway / Protected Routes
	api := r.Group("/api")
	api.Use(middleware.RateLimiterMiddleware(1, 5)) // 1 request per second, burst of 5
	api.Use(middleware.AuthMiddleware())
	{
		api.POST("/products", handlers.CreateProduct)
		api.GET("/products", handlers.ListProducts)
	}

	log.Println("Microservice starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
