package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/auth-service/internal/models"
	"github.com/yourusername/auth-service/internal/services"
	"github.com/sirupsen/logrus"
	"github.com/yourusername/auth-service/internal/utils"
)

type UserHandler struct {
	userService *services.UserService
	logger      *logrus.Logger
}

func NewUserHandler(userService *services.UserService, logger *logrus.Logger) *UserHandler {
	return &UserHandler{userService: userService, logger: logger}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		h.logger.Error("Failed to bind JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	password, err := utils.HashPassword(user.Password)

	if err != nil {
		h.logger.Error("Failed to create user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user", "msg": err})
	}

	user.Password = password

	if err := h.userService.CreateUser(&user); err != nil {
		h.logger.Error("Failed to create user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user", "msg": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user":    user,
	})
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		h.logger.Error("Failed to parse user ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := h.userService.GetUser(id)

	if err != nil {
		h.logger.Error("Failed to get user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User fetched successfully",
		"user":    user,
	})
} 

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		h.logger.Error("Failed to parse user ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		h.logger.Error("Failed to bind JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := h.userService.UpdateUser(id, &user); err != nil {
		h.logger.Error("Failed to update user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
		"user":    user,
	})
}

func (h *UserHandler) LoginUser (c *gin.Context){
	var user *models.UserLogin
	var users *models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		h.logger.Error("Failed to bind login user:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login user", "msg": err.Error()})
	}
	users, err = h.userService.LoginUser(user)
	if err != nil {
		h.logger.Error("Failed to login user:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login user", "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User fetched successfully",
		"user":    users,
	})
}