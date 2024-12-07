package main

import (
	"hermes/api"
	"hermes/config"
	"hermes/internal/repositories"
	"hermes/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializamos el router Gin
	router := gin.Default()

	// Configuración de la base de datos (usando GORM, por ejemplo)
	env := config.LoadEnv()
	db, err := config.Connection(env)
	if err != nil {
		panic("Error connecting to the database")
	}

	// Crear una instancia del repositorio y el servicio
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)

	// Configuración de las rutas y middleware
	api.SetupRouter(router, userService)

	// Iniciar el servidor en el puerto 8080
	router.Run(":8081")
}
