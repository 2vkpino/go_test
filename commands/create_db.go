package commands

import (
	"context"
	"fmt"
	"s3_file_uploader/config"
	"s3_file_uploader/utils"

	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
)

var createDbCmd = &cobra.Command{
	Use:   "create-db",
	Short: "Create MongoDB database and collections",
	Run: func(cmd *cobra.Command, args []string) {
		// Инициализация MongoDB
		utils.InitMongoDB()

		// Проверка на существование базы данных
		if databaseExists(config.MongoDBName) {
			utils.LogWarning(fmt.Sprintf("Database %s already exists", config.MongoDBName))
		} else {
			// Создание коллекции services
			createCollection("services")
		}
	},
}

func init() {
	rootCmd.AddCommand(createDbCmd)
}

func databaseExists(dbName string) bool {
	client := utils.MongoClient
	databases, err := client.ListDatabaseNames(context.TODO(), bson.D{})
	if err != nil {
		utils.LogError("Failed to list databases: ", err)
		return false
	}

	for _, d := range databases {
		if d == dbName {
			return true
		}
	}

	return false
}

func createCollection(collectionName string) {
	db := utils.MongoClient.Database(config.MongoDBName)
	collections, err := db.ListCollectionNames(context.TODO(), bson.D{})
	if err != nil {
		utils.LogError("Failed to list collections: ", err)
		return
	}

	for _, c := range collections {
		if c == collectionName {
			utils.LogInfo(fmt.Sprintf("Collection %s already exists", collectionName))
			return
		}
	}

	err = db.CreateCollection(context.TODO(), collectionName)
	if err != nil {
		utils.LogError(fmt.Sprintf("Failed to create collection %s: %v", collectionName, err))
		return
	}

	utils.LogInfo(fmt.Sprintf("Collection %s created successfully", collectionName))
}
