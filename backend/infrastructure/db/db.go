package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Montekrysto/ecomm/domain/entity"
)

func InitDB() (*gorm.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	connStr := fmt.Sprintf("CONNECTION STRING: user=%s password=%s dbname=%s port=%s host=%s sslmode=disable", dbUser, dbPassword, dbName, dbPort, dbHost)
	log.Printf("Connection %s", connStr)
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s host=%s sslmode=disable", dbUser, dbPassword, dbName, dbPort, dbHost)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	fmt.Println("MIGRATING DB")
	// Migrate the schema
	db.AutoMigrate(
		&entity.User{},
		&entity.Role{},
		&entity.Campaign{},
		&entity.Item{},
		&entity.Media{},
		&entity.Message{},
		&entity.Notification{},
	)

	// Create UUID extension
	// db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	return db, nil
}
