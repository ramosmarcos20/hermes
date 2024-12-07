package controllers

import (
	"hermes/internal/models/entities"
	"hermes/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	services *services.UserService
}

// Constructor para UserController
func NewUserController(services *services.UserService) *UserController {
	return &UserController{services: services}
}

func (c *UserController) GetUserProfile(ctx *gin.Context) {
	// Obtener el userID del contexto (que fue agregado por el middleware AuthMiddleware)
	userID, _ := ctx.Get("userID")

	// Usamos el userService para obtener el usuario por su ID
	user, err := c.services.GetUserById(userID.(uint)) // Uso de c.services en lugar de pasar userService
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Retornar los datos del usuario (puedes personalizar qué campos devolver)
	ctx.JSON(http.StatusOK, gin.H{
		"id":       user.ID,
		"email":    user.Email,
		"username": user.UserName,
	})
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.services.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving users"})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// Crear un nuevo usuario
func (c *UserController) CreateUser(ctx *gin.Context) {
	var user entities.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	errs := c.services.CreateUser(&user)
	if len(errs) > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": errs})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

// Obtener un usuario por ID
func (c *UserController) GetUserById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := c.services.GetUserById(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// Obtener un usuario por email
func (c *UserController) GetUserByEmail(ctx *gin.Context) {
	email := ctx.Query("email")
	if email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		return
	}

	user, err := c.services.GetUserByEmail(email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// Actualizar un usuario
func (c *UserController) UpdateUser(ctx *gin.Context) {
	var user entities.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	err := c.services.UpdateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// Eliminar un usuario
func (c *UserController) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = c.services.DeleteUser(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting user"})
		return
	}

	ctx.Status(http.StatusNoContent)
}
