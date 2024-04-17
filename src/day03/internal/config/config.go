package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	EsClientAddress   string
	MappingSchemaFile string
	DbFile            string
	UserName          string
	UserPassword      string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")

	if err != nil {
		return nil, err
	}

	esClientAddress := os.Getenv("ES_CLIENT_ADDRESS")

	if esClientAddress == "" {
		return nil, errors.New("ES_CLIENT_ADDRESS is not found")
	}

	mappingSchemaFile := os.Getenv("MAPPING_SCHEMA_FILE")

	if mappingSchemaFile == "" {
		return nil, errors.New("MAPPING_SCHEMA_FILE is not found")
	}

	dbFile := os.Getenv("DB_FILE")

	if dbFile == "" {
		return nil, errors.New("DB_FILE is not found")
	}

	userName := os.Getenv("USER_NAME")

	if userName == "" {
		return nil, errors.New("USER_NAME is not found")
	}

	userPassword := os.Getenv("USER_PASSWORD")

	if userPassword == "" {
		return nil, errors.New("USER_PASSWORD is not found")
	}

	return &Config{
		EsClientAddress:   esClientAddress,
		MappingSchemaFile: mappingSchemaFile,
		DbFile:            dbFile,
		UserName:          userName,
		UserPassword:      userPassword,
	}, nil
}
