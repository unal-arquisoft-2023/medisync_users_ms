package routes

import (
	"medysinc_user_ms/controllers"

	"github.com/labstack/echo/v4"
)

func PatientRoute(e *echo.Echo) {
	//pacientes

	e.POST("/patient", controllers.CreatePatient)
	e.GET("/patient/:patientId", controllers.GetPatient)
	e.PUT("/patient/:patientId", controllers.UpdatePatient)
	e.GET("/patients", controllers.GetAllPatients)
}
