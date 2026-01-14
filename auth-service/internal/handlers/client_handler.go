package handlers

import (
	"net/http"
	"strconv"

	// "github.com/gin-gonic/gin"
	// "github.com/yourusername/auth-service/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yourusername/auth-service/internal/models"
	"github.com/yourusername/auth-service/internal/services"
	// "github.com/yourusername/auth-service/internal/utils"
)

type ClientHandler struct {
	clientService *services.ClientService
	logger      *logrus.Logger
}

func NewClientHandler(clientService *services.ClientService, logger *logrus.Logger) *ClientHandler{
	return &ClientHandler{clientService: clientService, logger: logger}
	
}

func (h *ClientHandler) CreateClient(c *gin.Context){
	var client models.Client
	print("Creating client")
	
	err := c.ShouldBindJSON(&client)

	if err != nil {
		h.logger.Error("Failed to bind JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "msg": err.Error()})
		return
	}
	err = h.clientService.CreateClient(&client)

	if err != nil {
		h.logger.Error("Failed to create client:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create client", "msg": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Client created successfully", "id": client.ID})
}

func (h *ClientHandler) CreateClientFreelancer(c *gin.Context) {
	// This function is not implemented yet
	var clientFreelancer models.ClientFreelancer
	if err := c.ShouldBindJSON(&clientFreelancer); err != nil {
		h.logger.Error("Failed to bind JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "msg": err.Error()})
		return
	}

	err := h.clientService.CreateClientFreelancer(&clientFreelancer)
	if err != nil {
		h.logger.Error("Failed to create client-freelancer relationship:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create client-freelancer relationship", "msg": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Client-Freelancer relationship created successfully", "client_id": clientFreelancer.ClientId, "freelancer_id": clientFreelancer.FreelancerId})
}

func (h *ClientHandler) CreateClientAgency(c *gin.Context) {
	var clientAgency models.ClientAgency
	if err := c.ShouldBindJSON(&clientAgency); err != nil {
		h.logger.Error("Failed to bind JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "msg": err.Error()})
		return
	}

	err := h.clientService.CreateClientAgency(&clientAgency)
	if err != nil {
		h.logger.Error("Failed to create client-agency relationship:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create client-agency relationship", "msg": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Client-Agency relationship created successfully", "client_id": clientAgency.ClientId, "agency_id": clientAgency.AgencyId})
}


func (h *ClientHandler) CreateClientProject(c *gin.Context) {
	var clientProject models.ClientProject
	if err := c.ShouldBindJSON(&clientProject); err != nil {
		h.logger.Error("Failed to bind JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "msg": err.Error()})
		return
	}

	err := h.clientService.CreateClientProject(&clientProject)
	if err != nil {
		h.logger.Error("Failed to create client-project relationship:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create client-project relationship", "msg": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Client-Project relationship created successfully", "client_id": clientProject.ClientId, "project_id": clientProject.ProjectId})
}

func (h *ClientHandler) GetClientProjects(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		h.logger.Error("Failed to parse client ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid client ID"})
		return
	}
	projects, err := h.clientService.GetProjectsByClientId(uint(id))
	if err != nil {
		h.logger.Error("Failed to get projects by client ID:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get projects by client ID", "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"projects": projects})
}