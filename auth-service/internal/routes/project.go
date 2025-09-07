package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yourusername/auth-service/internal/handlers"
	"github.com/yourusername/auth-service/internal/services"
	"github.com/yourusername/auth-service/internal/utils"
)

type ProjectRoutes struct {
	router            *gin.Engine
	logger            *logrus.Logger
	projectService *services.ProjectService
}

func NewSetupProjectRoutes(router *gin.Engine, logger *logrus.Logger, projectService *services.ProjectService) *ProjectRoutes {
	return &ProjectRoutes{
		router:            router,
		logger:            logger,
		projectService: projectService,
	}
}

func (r *ProjectRoutes) SetupProjectRoutes() {
	handler := handlers.NewProjectHandler(r.projectService, r.logger)
	authMiddleware := utils.AuthMiddleware()

	r.router.POST("/projects", authMiddleware, handler.CreateProject)
}
