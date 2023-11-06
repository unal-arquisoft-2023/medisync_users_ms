package routes

import (
	. "medysinc_user_ms/controllers"
	. "medysinc_user_ms/controllers/patient"

	"github.com/labstack/echo/v4"
)

func PatientRoute(e *echo.Echo, patCon *PatientController) {
	e.POST("/patient", bindReq[CreatePatientRequest](patCon.CreatePatient))
	e.GET("/patient/:id", bindReq[UserIdRequest](patCon.GetPatient))
	e.PUT("/patient/:id", bindReq[UpdatePatientRequest](patCon.UpdatePatient))
	e.PUT("/patient/:id/suspend", bindReq[UserIdRequest](patCon.SuspendPatient))
	e.PUT("/patient/:id/activate", bindReq[UserIdRequest](patCon.ActivatePatient))
}
