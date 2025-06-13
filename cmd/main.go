package main

import (
	"prueba-go/config"
	"prueba-go/controllers"
	"prueba-go/models"
	"prueba-go/repository"
	"prueba-go/routes"
	"prueba-go/services"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	config.DB.AutoMigrate(&models.User{})

	userRepository := repository.NewUserRepository()
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	router := gin.Default()

	routes.RegisterRoutes(router, userController)

	router.Run(":8080")
}
