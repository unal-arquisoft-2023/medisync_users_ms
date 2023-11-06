package routes

import (
	c "medysinc_user_ms/controllers"
	pc "medysinc_user_ms/controllers/patient"

	"github.com/labstack/echo/v4"
)

func PatientRoute(e *echo.Echo, patCon *pc.PatientController) {
	e.POST("/patient", bindAndValidateReq[pc.CreatePatientRequest](patCon.CreatePatient))
	e.GET("/patient/:id", bindAndValidateReq[c.UserIdRequest](patCon.GetPatient))
	e.PUT("/patient/:id", bindAndValidateReq[pc.UpdatePatientRequest](patCon.UpdatePatient))
	e.PUT("/patient/:id/suspend", bindAndValidateReq[c.UserIdRequest](patCon.SuspendPatient))
	e.PUT("/patient/:id/activate", bindAndValidateReq[c.UserIdRequest](patCon.ActivatePatient))
}
