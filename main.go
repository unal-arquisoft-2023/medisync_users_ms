package main

import (
	"medysinc_user_ms/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	//routes
	routes.StaffRoute(e)
	routes.DoctorRoute(e)
	routes.PatientRoute(e)

	e.Logger.Fatal(e.Start(":6000"))
}
