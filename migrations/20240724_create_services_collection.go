package migrations

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"s3_file_uploader/utils"
)

func Up() error {
	if utils.MongoClient == nil {
		return fmt.Errorf("MongoClient is not initialized")
	}

	db := utils.MongoClient.Database(utils.MongoDBName)

	// Проверяем, существует ли коллекция
	collections, err := db.ListCollectionNames(context.TODO(), bson.D{})
	if err != nil {
		return err
	}

	for _, name := range collections {
		if name == "services" {
			utils.MongoLogInfo("Collection 'services' already exists")
			return nil
		}
	}

	// Создаем коллекцию
	err = db.CreateCollection(context.TODO(), "services", createCollectionOptions())
	if err != nil {
		return err
	}

	utils.MongoLogInfo("Collection 'services' created successfully")
	return nil
}

func Down() error {
	if utils.MongoClient == nil {
		return fmt.Errorf("MongoClient is not initialized")
	}

	db := utils.MongoClient.Database(utils.MongoDBName)

	// Удаляем коллекцию
	err := db.Collection("services").Drop(context.TODO())
	if err != nil {
		return err
	}

	utils.MongoLogInfo("Collection 'services' dropped successfully")
	return nil
}

func createCollectionOptions() *mongo.CollectionOptions {
	return &mongo.CollectionOptions{
		Validator: bson.D{
			{"$jsonSchema", bson.D{
				{"bsonType", "object"},
				{"required", []string{"name", "nfs_dir", "s3_bucket"}},
				{"properties", bson.D{
					{"name", bson.D{
						{"bsonType", "string"},
						{"description", "Service name is required and must be a string"},
					}},
					{"nfs_dir", bson.D{
						{"bsonType", "string"},
						{"description", "NFS directory is required and must be a string"},
					}},
					{"s3_bucket", bson.D{
						{"bsonType", "string"},
						{"description", "S3 bucket name is required and must be a string"},
					}},
					{"s3_region", bson.D{
						{"bsonType", "string"},
						{"description", "S3 region is optional and must be a string if provided"},
					}},
					{"created_at", bson.D{
						{"bsonType", []string{"date", "null"}},
						{"description", "Creation date must be a valid date or null"},
					}},
					{"updated_at", bson.D{
						{"bsonType", []string{"date", "null"}},
						{"description", "Last update date must be a valid date or null"},
					}},
				}},
			}},
		},
	}
}
