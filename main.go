package main

import (
	"context"
	"fmt"
	con "medysinc_user_ms/controllers"
	pcon "medysinc_user_ms/controllers/patient"
	"medysinc_user_ms/resources/configuration"
	mdb "medysinc_user_ms/resources/database/mongo_database"
	mongoRepos "medysinc_user_ms/resources/users/mongodb"
	"medysinc_user_ms/resources/validation"
	"medysinc_user_ms/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Hello, world!")

	ctx := context.Background()

	config, err := configuration.NewConfigurationGodotEnv(".env")
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

	e.Logger.Fatal(e.Start(":6000"))
}
