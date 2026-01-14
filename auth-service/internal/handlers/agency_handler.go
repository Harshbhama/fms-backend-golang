package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yourusername/auth-service/internal/models"
	"github.com/yourusername/auth-service/internal/services"
	"net/http"
)

type AgencyHandler struct {
	Service *services.AgencyService
	logger  *logrus.Logger
}

func NewAgencyHandler(service *services.AgencyService, logger *logrus.Logger) *AgencyHandler {
	return &AgencyHandler{Service: service, logger: logger}
}

func (h *AgencyHandler) CreateAgency(c *gin.Context) {
	var agency models.Agency

	err := c.ShouldBindJSON(&agency)
	if err != nil {
		h.logger.Error("Failed to bind JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "msg": err.Error()})
		return
	}

	err = h.Service.CreateAgency(&agency)
	if err != nil {
		h.logger.Error("Failed to create agency:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create agency", "msg": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Agency created successfully", "id": agency.ID})
}
