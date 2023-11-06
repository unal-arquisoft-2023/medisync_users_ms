package main

import (
	"context"
	con "medysinc_user_ms/controllers"
	pcon "medysinc_user_ms/controllers/patient"
	"medysinc_user_ms/resources/configuration"
	mdb "medysinc_user_ms/resources/database/mongo_database"
	mongoRepos "medysinc_user_ms/resources/users/mongodb"
	"medysinc_user_ms/resources/validation"
	"medysinc_user_ms/routes"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	ctx := context.Background()

	config, err := configuration.NewConfigurationGodotEnv(".env")
	if err != nil {
		panic(err)
	}

	port, err := config.Get("PORT")
	if err != nil {
		panic(err)
	}

	db := mdb.NewMongoDatabase(ctx, config)

	e := echo.New()
	val := validation.NewMedisyncValidator()
	pcon.AddCustomDTOValidations(val)
	con.AddCustomDTOValidations(val)

	e.Validator = val

	patRepo := mongoRepos.NewMongoPatientRepository(db.PatCol)
	patCon := pcon.NewPatientController(patRepo)

	routes.PatientRoute(e, patCon)

	// Start server
	go func() {
		if err := e.Start(":" + port); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
