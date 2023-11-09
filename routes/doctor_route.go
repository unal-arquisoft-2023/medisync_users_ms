package routes

import (
	c "medysinc_user_ms/controllers"
	dc "medysinc_user_ms/controllers/doctor"

	"github.com/labstack/echo/v4"
)

func DoctorRoute(e *echo.Echo, docCon *dc.DoctorController) {
	e.POST("/doctor", bindAndValidateReq[dc.CreateDoctorRequest](docCon.CreateDoctor))
	e.GET("/doctor/:id", bindAndValidateReq[c.UserIdRequest](docCon.GetDoctor))
	e.PUT("/doctor/:id", bindAndValidateReq[dc.UpdateDoctorRequest](docCon.UpdateDoctor))
	e.PUT("/doctor/:id/suspend", bindAndValidateReq[c.UserIdRequest](docCon.SuspendDoctor))
	e.PUT("/doctor/:id/activate", bindAndValidateReq[c.UserIdRequest](docCon.ActivateDoctor))
	e.GET("/doctors", docCon.GetAllDoctors)
	e.GET("/doctors/:specialty", bindAndValidateReq[c.SpecialtyRequest](docCon.GetAllDoctorsBySpecialty))
}
