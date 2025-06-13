package routes

import (
	"prueba-go/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, userController *controllers.UserController) {
	v1 := router.Group("/api")
	{
		v1.POST("/users", userController.CreateUser)
		v1.GET("/users", userController.GetUsers)
		v1.GET("/users/:id", userController.GetUserByID)
		v1.PUT("/users/:id", userController.UpdateUser)
		v1.DELETE("/users/:id", userController.DeleteUser)
	}
}