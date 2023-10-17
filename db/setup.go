package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

// ConnectDB conecta a la base de datos MongoDB y devuelve un cliente MongoDB.
func ConnectDB() (*mongo.Client, *BaseCollections) {
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

	coll := BaseCollections{
		Users:   GetCollection(client, "users"),
		Staff:   GetCollection(client, "staff"),
		Doctor:  GetCollection(client, "doctor"),
		Patient: GetCollection(client, "patient"),
	}

	fmt.Println("Connected to MongoDB")
	return client, &coll
}

// Client instance
var DB, Collections = ConnectDB()
