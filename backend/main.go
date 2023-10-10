package main

import (
	"github.com/Montekrysto/ecomm/app/usecase"
	"github.com/Montekrysto/ecomm/infrastructure/api/handler"
	"github.com/Montekrysto/ecomm/infrastructure/api/routes"
	"github.com/Montekrysto/ecomm/infrastructure/db"
	"github.com/gin-gonic/gin"
)

func main() {
	// if err := godotenv.Load("../.env"); err != nil {
	// 	log.Fatalf("Error loading .env file in mainasdf: %v", err)
	// }

	// Initialize the database connection
	dbInstance, err := db.InitDB()
	if err != nil {
		panic(err)
	}

	// Initialize repositories
	userRepo := db.NewUserRepository(dbInstance)

	// Initialize use cases
	userUseCase := usecase.NewUserUseCase(userRepo)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userUseCase)

	// Initialize the router
	router := gin.Default()

	// Setup the routes
	routes.SetupUserRoutes(router, userHandler)

	// Start the server
	if err := router.Run(); err != nil {
		panic(err)
	}
}
