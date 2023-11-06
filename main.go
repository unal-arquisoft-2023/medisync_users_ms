package main

import (
	"context"
	"fmt"
	pc "medysinc_user_ms/controllers/patient"
	"medysinc_user_ms/resources/configuration"
	mongoRepos "medysinc_user_ms/resources/users/mongodb"
	"medysinc_user_ms/routes"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	fmt.Println("Hello, world!")

	ctx := context.Background()

	configuration, err := configuration.NewConfigurationGodotEnv(".env")

	if err != nil {
		panic(err)
	}

	mongoURI, err := configuration.Get("MONGO_URI")

	fmt.Println("MONGO_URI")
	fmt.Println(mongoURI)
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
	db := mongoClient.Database("medisync")
	patColl := db.Collection("patients")
	patRepo := mongoRepos.NewMongoPatientRepository(patColl)

	e := echo.New()
	patCon := pc.NewPatientController(patRepo)

	routes.PatientRoute(e, patCon)

	e.Logger.Fatal(e.Start(":6000"))
}
