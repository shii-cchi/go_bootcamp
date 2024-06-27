package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type Record struct {
	gorm.Model
	SessionId string
	Frequency float64
	Timestamp time.Time
}

func ConnectToDb() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5433 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&Record{})

	if err != nil {
		return nil, fmt.Errorf("failed to automigrate: %v", err)
	}

	return db, nil
}
