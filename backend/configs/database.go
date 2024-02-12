package configs

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
	SLLMode  string
}

func LoadEnv(path string) error {
	return godotenv.Load(path)
}

func buildConnectionString(data DatabaseConfig) string {
	return fmt.Sprintf("host=%s user=%s port=%s password=%s dbname=%s sslmode=%s", data.Host, data.User, data.Port, data.Password, data.Name, data.SLLMode)
}

func SetUpDatabase() (*gorm.DB, error) {
	if envError := LoadEnv("../.env"); envError != nil {
		return nil, envError
	}

	config := DatabaseConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		User:     os.Getenv("POSTGRES_USER"),
		Name:     os.Getenv("POSTGRES_NAME"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Port:     os.Getenv("POSTGRES_PORT"),
		SLLMode:  os.Getenv("POSTGRES_SLLMODE"),
	}

	database, connectionError := gorm.Open(postgres.Open(buildConnectionString(config)), &gorm.Config{})

	if connectionError != nil {
		return nil, connectionError
	}

	return database, nil
}
