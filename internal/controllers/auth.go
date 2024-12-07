package controllers

import (
	"errors"
	"hermes/internal/models/entities"
	"hermes/internal/services"
	"hermes/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginInput struct {
	Email    string `json:"email"`
	UserName string `json:"user_name"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context, userService *services.UserService) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	var user *entities.User
	var err error

	// Primero intentamos buscar por email
	if input.Email != "" {
		user, err = userService.GetUserByEmail(input.Email)
	} else if input.UserName != "" {
		// Si no hay email, intentamos buscar por username
		user, err = userService.GetUserByUsername(input.UserName)
	}

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or username or password"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	// Verificamos la contraseña
	if !utils.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or username or password"})
		return
	}

	// Generamos el token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	// Respondemos con el token
	c.JSON(http.StatusOK, gin.H{"token": token})
}
