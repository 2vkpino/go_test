package migrations

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"s3_file_uploader/utils"
)

func Up() error {
	db := utils.MongoClient.Database("mydatabase")

	// Определение параметров валидации для коллекции
	collectionOptions := options.CreateCollection().SetValidator(bson.D{
		{"$jsonSchema", bson.D{
			{"bsonType", "object"},
			{"required", bson.A{"name", "nfs_dir", "s3_bucket"}},
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
					{"bsonType", bson.A{"date", "null"}},
					{"description", "Creation date must be a valid date or null"},
				}},
				{"updated_at", bson.D{
					{"bsonType", bson.A{"date", "null"}},
					{"description", "Last update date must be a valid date or null"},
				}},
			}},
		}},
	})

	// Проверка на существование коллекции
	collections, err := db.ListCollectionNames(context.TODO(), bson.D{})
	if err != nil {
		return err
	}

	for _, name := range collections {
		if name == "services" {
			utils.LogInfo("Collection 'services' already exists")
			return nil
		}
	}

	// Создание коллекции с валидацией
	err = db.CreateCollection(context.TODO(), "services", collectionOptions)
	if err != nil {
		return err
	}

	utils.LogInfo("Collection 'services' created successfully with validation")
	return nil
}

func Down() error {
	db := utils.MongoClient.Database("mydatabase")

	// Удаление коллекции
	err := db.Collection("services").Drop(context.TODO())
	if err != nil {
		return err
	}

	utils.LogInfo("Collection 'services' dropped successfully")
	return nil
}
