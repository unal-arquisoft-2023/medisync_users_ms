package routes

import (
	"medysinc_user_ms/controllers/patient"

	"github.com/labstack/echo/v4"
)

func PatientRoute(e *echo.Echo, patCon *patient.PatientController) {
	//pacientes
	e.POST("/patient", bindJSON[patient.CreatePatientRequest](patCon.CreatePatient))
}
