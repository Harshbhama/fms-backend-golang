package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yourusername/auth-service/internal/handlers"
	"github.com/yourusername/auth-service/internal/services"
)

type AgencyRoutes struct {
	router        *gin.Engine
	logger        *logrus.Logger
	agencyService *services.AgencyService
}

func NewSetupAgencyRoutes(router *gin.Engine, logger *logrus.Logger, agencyService *services.AgencyService) *AgencyRoutes {
	return &AgencyRoutes{
		router:        router,
		logger:        logger,
		agencyService: agencyService,
	}
}

func (r *AgencyRoutes) SetupAgency() {
	handler := handlers.NewAgencyHandler(r.agencyService, r.logger)

	r.router.POST("/agency", handler.CreateAgency)
}
