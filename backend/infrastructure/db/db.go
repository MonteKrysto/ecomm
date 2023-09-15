package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Montekrysto/ecomm/domain/entity"
)

func InitDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", "user", "password", "dbname", "5432")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

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
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	return db, nil
}
