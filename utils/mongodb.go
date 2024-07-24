package utils

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"s3_file_uploader/config"
)

// MongoClient представляет подключение к MongoDB.
var MongoClient *mongo.Client

// ServiceCollection представляет коллекцию "services" в MongoDB.
var ServiceCollection *mongo.Collection

// InitMongoDB инициализирует подключение к MongoDB и настраивает коллекцию.
func InitMongoDB() {
	clientOptions := options.Client().ApplyURI(config.MongoURI)
	var err error
	MongoClient, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		MongoLogError("Failed to connect to MongoDB: ", err)
		return
	}

	err = MongoClient.Ping(context.Background(), nil)
	if err != nil {
		MongoLogError("Failed to ping MongoDB: ", err)
		return
	}

	// Инициализация коллекции "services".
	ServiceCollection = MongoClient.Database(config.MongoDBName).Collection("services")
	MongoLogInfo("MongoDB connected and collection 'services' initialized")
}

// MongoLogInfo выводит информационное сообщение.
func MongoLogInfo(message string) {
	log.Printf("INFO: %s", message)
}

// MongoLogError выводит сообщение об ошибке.
func MongoLogError(message string, err error) {
	log.Printf("ERROR: %s %v", message, err)
}
