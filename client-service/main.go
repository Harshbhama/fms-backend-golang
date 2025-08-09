package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client
 
func main() {
	// Initialize Redis
	initRedis()

	r := gin.Default()

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
