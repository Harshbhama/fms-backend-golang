package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yourusername/auth-service/internal/models"
	"github.com/yourusername/auth-service/internal/services"
)

type ProjectHandler struct {
	Service *services.ProjectService
	logger  *logrus.Logger
}

func NewProjectHandler(service *services.ProjectService, logger *logrus.Logger) *ProjectHandler {
	return &ProjectHandler{Service: service, logger: logger}
}

func (h *ProjectHandler) CreateProject(c *gin.Context) {
	var project models.Projects
	h.logger.Info("Creating project")

	err := c.ShouldBindJSON(&project)
	if err != nil {
		h.logger.Error("Failed to bind JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "msg": err.Error()})
		return
	}

	err = h.Service.CreateProject(&project)
	if err != nil {
		h.logger.Error("Failed to create project:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create project", "msg": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Project created successfully", "id": project.ID})
}