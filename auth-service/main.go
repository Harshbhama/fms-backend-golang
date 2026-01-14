package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/yourusername/auth-service/internal/config"
	"github.com/yourusername/auth-service/internal/repositories"
	"github.com/yourusername/auth-service/internal/services"
	"github.com/yourusername/auth-service/internal/routes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

var db *gorm.DB
var redisClient *redis.Client

func main() {
	// Initialize config (includes DB connection)
	cfg, err := config.NewConfig()
	logger := logrus.New()
	if err != nil {
		logger.Fatalf("failed to load config or connect to DB: %v", err)
	}

	authRepo := repositories.NewUserRepository(cfg.DB)
	authService := services.NewUserService(authRepo)

	clientRepo := repositories.NewClientRepository(cfg.DB)
	clientService := services.NewClientService(clientRepo)

	freelancerRepo := repositories.NewFreelancerRepository(cfg.DB)
	freelancerService := services.NewFreelancerService(freelancerRepo)

	agencyRepo := repositories.NewAgencyRepository(cfg.DB)
	agencyService := services.NewAgencyService(agencyRepo)

	projectRepo := repositories.NewProjectRepository(cfg.DB)
	projectService := services.NewProjectService(projectRepo)

	// Initialize Gin router
	router := gin.Default()

	// ✅ Enable CORS (allow all origins)
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true, // allow requests from anywhere
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:   []string{"Content-Length"},
		// ⚠️ Must be false when AllowAllOrigins = true
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	// Setup routes
	authRoutes := routes.NewSetupRoutes(router, logger, authService)
	clientRoutes := routes.NewSetupClientRoutes(router, logger, clientService)
	projectRoutes := routes.NewSetupProjectRoutes(router, logger, projectService)
	freelancerRoutes := routes.NewSetupFreelancerRoutes(router, logger, freelancerService)
	agencyRoutes := routes.NewSetupAgencyRoutes(router, logger, agencyService)

	authRoutes.Setup()
	clientRoutes.SetupClient()
	projectRoutes.SetupProjectRoutes()
	freelancerRoutes.SetupFreelancer()
	agencyRoutes.SetupAgency()

	// Start server
	port := ":8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = ":" + envPort
	}
	router.Run(port)
}

func initDB() {
	dsn := os.Getenv("DB_CONN")
	if dsn == "" {
		dsn = "host=postgres user=admin password=admin123 dbname=auth_db port=5432 sslmode=disable"
	}

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate the schema
	db.AutoMigrate(&User{})
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

	// Test the connection with a 5-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}
}
