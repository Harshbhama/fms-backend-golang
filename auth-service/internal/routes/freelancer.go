package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yourusername/auth-service/internal/handlers"
	"github.com/yourusername/auth-service/internal/services"
)

type FreelancerRoutes struct {
	router            *gin.Engine
	logger            *logrus.Logger
	freelancerService *services.FreelancerService
}

func NewSetupFreelancerRoutes(router *gin.Engine, logger *logrus.Logger, freelancerService *services.FreelancerService) *FreelancerRoutes {
	return &FreelancerRoutes{
		router:            router,
		logger:            logger,
		freelancerService: freelancerService,
	}
}

func (r *FreelancerRoutes) SetupFreelancer() {
	handler := handlers.NewFreelancerHandler(r.freelancerService, r.logger)
	r.router.POST("/freelancers", handler.CreateFreelancer)
	r.router.POST("/freelancer-rates", handler.CreateFreelancerRates)
	r.router.POST("/freelancer-project", handler.CreateFreelancerProject)
	r.router.POST("/freelancer-timesheet", handler.CreateFreelancerTimesheet)
	r.router.POST("/freelancer-timesheet-metadata", handler.CreateFreelancerTimesheetMetadata)
}
