package routes

import (
	"github.com/Montekrysto/ecomm/infrastructure/api/handler"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine, userHandler *handler.UserHandler) *gin.Engine {
	// router := gin.Default()

	// Define your routes here
	router.GET("/users", userHandler.GetAllUsers)
	router.GET("/users/:id", userHandler.GetUser)
	router.POST("/users", userHandler.CreateUser)
	// router.PUT("/users/:id", userHandler.UpdateUser)
	// router.DELETE("/users/:id", userHandler.DeleteUser)

	return router
}
