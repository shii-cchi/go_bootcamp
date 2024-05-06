package config

import (
	"errors"
	"os"
)
import "github.com/joho/godotenv"

type Config struct {
	Port          string
	DbURI         string
	AdminLogin    string
	AdminPassword string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load("blog/.env")

	if err != nil {
		return nil, err
	}

	port := os.Getenv("PORT")

	if port == "" {
		return nil, errors.New("PORT is not found")
	}

	dbURI := os.Getenv("DB_URI")

	if dbURI == "" {
		return nil, errors.New("DB_URL is not found")
	}

	adminLogin := os.Getenv("ADMIN_LOGIN")

	if adminLogin == "" {
		return nil, errors.New("ADMIN_LOGIN is not found")
	}

	adminPassword := os.Getenv("ADMIN_PASSWORD")

	if adminPassword == "" {
		return nil, errors.New("ADMIN_PASSWORD is not found")
	}

	return &Config{
		Port:          port,
		DbURI:         dbURI,
		AdminLogin:    adminLogin,
		AdminPassword: adminPassword,
	}, nil
}
