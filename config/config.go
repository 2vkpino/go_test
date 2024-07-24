package config

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	MongoURI    string
	MongoDBName string
	LogFile     string
	LogLevel    int
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	MongoURI = os.Getenv("MONGO_URI")
	MongoDBName = os.Getenv("MONGO_DB_NAME")
	LogFile = os.Getenv("LOG_FILE")

	if LogFile == "" {
		// Установка пути по умолчанию для лог-файла
		projectRoot, err := os.Getwd()
		if err != nil {
			log.Fatalf("Failed to get current directory: %v", err)
		}
		LogFile = filepath.Join(projectRoot, "var", "log", "s3_file_uploader.log")
	}

	LogLevel, err = strconv.Atoi(os.Getenv("LOG_LEVEL"))
	if err != nil {
		LogLevel = 1 // Уровень по умолчанию: INFO
	}
}
