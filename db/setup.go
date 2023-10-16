package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDB conecta a la base de datos MongoDB y devuelve un cliente MongoDB.
func ConnectDB() *mongo.Client {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI(EnvMongoURI())

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")
	return client
}

// Client instance
var DB *mongo.Client = ConnectDB()

// GetCollection obtiene una colección de la base de datos.
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("golangAPI").Collection(collectionName)
	return collection
}

type BaseCollections struct {
	Users   *mongo.Collection
	Staff   *mongo.Collection
	Doctor  *mongo.Collection
	Patient *mongo.Collection
}

var Collections = BaseCollections{
	Users:   GetCollection(DB, "users"),
	Staff:   GetCollection(DB, "staff"),
	Doctor:  GetCollection(DB, "doctor"),
	Patient: GetCollection(DB, "patient"),
}
