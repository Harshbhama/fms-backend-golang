package routes

import (
	"github.com/gin-gonic/gin"
	// "github.com/harshbhama/go-gin-postgres-app/internal/config"
	"github.com/yourusername/auth-service/internal/handlers"
	// "github.com/harshbhama/go-gin-postgres-app/internal/repositories"
	"github.com/yourusername/auth-service/internal/services"
	"github.com/sirupsen/logrus"
)

// SetupRoutes sets up all the API routes
type SetupRoutes struct {
	Router *gin.Engine
	Logger *logrus.Logger
	UserService *services.UserService
}

func NewSetupRoutes(router *gin.Engine, logger *logrus.Logger, userService *services.UserService) *SetupRoutes {
	return &SetupRoutes{
		Router: router,
		Logger: logger,
		UserService: userService,
	}
}

func (r *SetupRoutes) Setup() {
	
	userHandler := handlers.NewUserHandler(r.UserService, r.Logger)

	// API Routes
	r.Router.GET("/health", r.healthCheck)
	r.Router.GET("/status", r.status)
	r.Router.POST("/login", userHandler.LoginUser)
	r.Router.POST("/users", userHandler.CreateUser)
	r.Router.GET("/users/:id", userHandler.GetUser)
	r.Router.PUT("/users/:id", userHandler.UpdateUser)

}

func (r *SetupRoutes) healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "healthy",
		"message": "Server is running",
	})
}

func (r *SetupRoutes) status(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "ok",
		"service": "go-gin-postgres-app",
	})
}

func (r *SetupRoutes) login(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "ok",
		"message": "Login successful",
	})
}
