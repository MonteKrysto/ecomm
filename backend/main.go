package main

import (
	"log"
	"os"
	"strconv"
	"github.com/Montekrysto/ecomm/infrastructure/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env file found")
	}
}

func main() {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))

	// Use these variables to initialize your database
	db, err := db.NewDatabase(dbUser, dbPassword, dbHost, dbName, dbPort)
	if err != nil {
		log.Fatalf("could not initialize database: %v", err)
	}
	defer db.Close()

	// Initialize Gin router
	router := gin.Default()

	// Register routes
	registerRoutes(router)

	// Start server
	router.Run(":8080")
}

func registerRoutes(router *gin.Engine) {
	// Your route registrations here, e.g.:
	// router.GET("/path", handlerFunc)
}
