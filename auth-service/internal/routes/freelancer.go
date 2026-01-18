package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yourusername/auth-service/internal/handlers"
	"github.com/yourusername/auth-service/internal/services"
	"github.com/yourusername/auth-service/internal/utils"
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
	authMiddleware := utils.AuthMiddleware()

	r.router.POST("/freelancers", handler.CreateFreelancer)
	r.router.POST("/freelancer-rates", authMiddleware, handler.CreateFreelancerRates)
	r.router.POST("/freelancer-project", authMiddleware, handler.CreateFreelancerProject)
	r.router.POST("/freelancer-timesheet", authMiddleware, handler.CreateFreelancerTimesheet)
	r.router.POST("/freelancer-timesheet-metadata", authMiddleware, handler.CreateFreelancerTimesheetMetadata)
	r.router.GET("/freelancer-by-client/:id", authMiddleware, handler.GetFreelancerForClient)
	r.router.GET("/projects-by-freelancer/:id", authMiddleware, handler.GetProjectsForFreelancer)
	r.router.GET("/clients-by-freelancer/:id", authMiddleware, handler.GetClientsForFreelancer)
}
