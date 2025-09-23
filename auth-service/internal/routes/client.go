package routes

import (
	"github.com/gin-gonic/gin"
	// "github.com/harshbhama/go-gin-postgres-app/internal/config"
	"github.com/yourusername/auth-service/internal/handlers"
	// "github.com/harshbhama/go-gin-postgres-app/internal/repositories"
	"github.com/yourusername/auth-service/internal/services"
	"github.com/sirupsen/logrus"
	"github.com/yourusername/auth-service/internal/utils"
)

// SetupRoutes sets up all the API routes
type SetupClientRoutes struct {
	Router *gin.Engine
	Logger *logrus.Logger
	ClientService *services.ClientService
}

func NewSetupClientRoutes(router *gin.Engine, logger *logrus.Logger, clientService *services.ClientService) *SetupClientRoutes {
	return &SetupClientRoutes{
		Router: router,
		Logger: logger,
		ClientService: clientService,
	}
}

func (r *SetupClientRoutes) SetupClient() {
	
	clientHandler := handlers.NewClientHandler(r.ClientService, r.Logger)
	authMiddleware := utils.AuthMiddleware()

	r.Router.POST("/client", clientHandler.CreateClient)
	r.Router.POST("/client-freelancer", clientHandler.CreateClientFreelancer)
	r.Router.POST("/client-project", authMiddleware, clientHandler.CreateClientProject)
}



