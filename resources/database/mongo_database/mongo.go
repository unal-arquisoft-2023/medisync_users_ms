package mongodatabase

import (
	"context"
	"medysinc_user_ms/resources/configuration"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDatabase struct {
	PatCol *mongo.Collection
}

func NewMongoDatabase(ctx context.Context, config configuration.ConfigurationRepository) *MongoDatabase {

	mongoURI, err := config.Get("MONGO_URI")
	if err != nil {
		panic(err)
	}
	mongoDB, err := config.Get("MONGO_DB")
	if err != nil {
		panic(err)
	}

	clientOptions :=
		options.Client().
			ApplyURI(mongoURI)

	mongoClient, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		panic(err)
	}

	db := mongoClient.Database(mongoDB)
	patColl := db.Collection("patients")

	return &MongoDatabase{
		PatCol: patColl,
	}
}
