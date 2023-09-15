package db

import (
	"fmt"
	"time"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(user, password, host, dbname string, port int) (*gorm.DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s host=%s dbname=%s port=%d sslmode=disable", user, password, host, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, err := db.DB()
	if err != nil {
			panic("failed to get database instance")
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	err = db.AutoMigrate(
		&User{},
		&CreatorProfile{},
		&ContributorProfile{},
		&CustomerProfile{},
		&Campaign{},
		&Item{},
		&Photo{},
		&Video{},
		&Message{},
		&Notification{},
		&Invitation{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
