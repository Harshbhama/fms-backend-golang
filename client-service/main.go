package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func main() {
	// Initialize Redis
	initRedis()

	r := gin.Default()

	// Enable CORS
	r.Use(cors.New(cors.Config{
    AllowAllOrigins:  true, // 👈 this enables all origins
    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: false, // ⚠️ must be false when AllowAllOrigins is true
    MaxAge:           12 * time.Hour,
	}))

	// Routes
	r.GET("/protected", authMiddleware(), protectedHandler)
	r.GET("/health", healthCheck)

	port := ":8081"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = ":" + envPort
	}

	r.Run(port)
}

func initRedis() {
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "redis:6379"
	}

	redisPassword := os.Getenv("REDIS_PASSWORD")

	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       0,
	})

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Implementation for authentication middleware
		// This would verify the JWT token from the request
		c.Next()
	}
}

func protectedHandler(c *gin.Context) {
	// This is a protected route that requires authentication
	c.JSON(http.StatusOK, gin.H{"message": "This is a protected endpoint"})
}
