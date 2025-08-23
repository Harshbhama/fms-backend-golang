package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yourusername/auth-service/internal/models"
	"github.com/yourusername/auth-service/internal/services"
)

type FreelancerHandler struct {
	Service *services.FreelancerService
	logger  *logrus.Logger
}

func NewFreelancerHandler(service *services.FreelancerService, logger *logrus.Logger) *FreelancerHandler {
	return &FreelancerHandler{Service: service, logger: logger}
}

func (h *FreelancerHandler) CreateFreelancer(c *gin.Context) {
	var freelancer models.Freelancer
	print("Creating freelancer")

	err := c.ShouldBindJSON(&freelancer)
	if err != nil {
		h.logger.Error("Failed to bind JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "msg": err.Error()})
		return
	}

	err = h.Service.CreateFreelancer(&freelancer)
	if err != nil {
		h.logger.Error("Failed to create freelancer:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create freelancer", "msg": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Freelancer created successfully", "id": freelancer.ID})
}

func (h *FreelancerHandler) CreateFreelancerRates(c *gin.Context) {
	var freelancerRates models.FreelancerRates
	err := c.ShouldBindJSON(&freelancerRates)
	if err != nil {
		h.logger.Error("Failed to bind JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "msg": err.Error()})
		return
	}

	err = h.Service.CreateFreelancerRates(&freelancerRates)
	if err != nil {
		h.logger.Error("Failed to create freelancer rates:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create freelancer rates", "msg": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Freelancer rates created successfully", "id": freelancerRates.ID})
}

func (h *FreelancerHandler) CreateFreelancerProject(c *gin.Context) {
	var freelancerProject models.FreelancerProject
	if err := c.ShouldBindJSON(&freelancerProject); err != nil {
		h.logger.Error("Failed to bind JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "msg": err.Error()})
		return
	}

	err := h.Service.CreateFreelancerProject(&freelancerProject)
	if err != nil {
		h.logger.Error("Failed to create freelancer project:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create freelancer project", "msg": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Freelancer project created successfully", "freelancer_id": freelancerProject.FreelancerId, "project_id": freelancerProject.ProjectId})
}