package handlers

import (
	"net/http"
	"strconv"
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

func (h *FreelancerHandler) CreateFreelancerTimesheet(c *gin.Context) {
	var timesheet models.FreelancerTimesheet
	if err := c.ShouldBindJSON(&timesheet); err != nil {
		h.logger.Error("Failed to bind JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "msg": err.Error()})
		return
	}
	err := h.Service.CreateFreelancerTimesheet(&timesheet)
	if err != nil {
		h.logger.Error("Failed to create freelancer timesheet:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create freelancer timesheet", "msg": err.Error()})
		return
	}
	// Here you would typically call a service method to handle the timesheet creation.
	// For demonstration, we'll just log and return the received data.

	h.logger.Info("Received timesheet:", timesheet)

	c.JSON(http.StatusCreated, gin.H{"message": "Freelancer timesheet created successfully", "freelancer_id": timesheet.FreelancerID, "project_id": timesheet.ProjectID})
}

func (h *FreelancerHandler) CreateFreelancerTimesheetMetadata(c *gin.Context) {
	var metadata models.FreelancerTimesheetMetadata
	if err := c.ShouldBindJSON(&metadata); err != nil {
		h.logger.Error("Failed to bind JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "msg": err.Error()})
		return
	}

	err := h.Service.CreateFreelancerTimesheetMetadata(&metadata)
	if err != nil {
		h.logger.Error("Failed to create freelancer timesheet metadata:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create freelancer timesheet metadata", "msg": err.Error()})
		return
	}
	h.logger.Info("Received timesheet metadata:", metadata)

	c.JSON(http.StatusCreated, gin.H{"message": "Freelancer timesheet metadata created successfully", "metadata_id": metadata.MetadataID, "timesheet_id": metadata.TimesheetID})
}
func (h *FreelancerHandler) GetFreelancerForClient(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		h.logger.Error("Failed to parse client ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid client ID"})
		return
	}
	freelancers, err := h.Service.GetFreelancerByClientID(uint(id))
	h.logger.Info("Freelancers fetched for client ID:", id, freelancers)
	if err != nil {
		print("Failed to get freelancers for client--------------------------------", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get freelancers for client", "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "Freelancers fetched successfully12",
		"freelancers": freelancers,
	})
}