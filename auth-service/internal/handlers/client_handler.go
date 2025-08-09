package handlers

import (
	"net/http"
	// "strconv"

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

