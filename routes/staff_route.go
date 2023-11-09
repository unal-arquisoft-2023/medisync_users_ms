package routes

import (
	c "medysinc_user_ms/controllers"
	sc "medysinc_user_ms/controllers/staff"

	"github.com/labstack/echo/v4"
)

func StaffRoute(e *echo.Echo, staffCon *sc.StaffController) {
	e.POST("/staff", bindAndValidateReq[sc.CreateStaffRequest](staffCon.CreateStaff))
	e.GET("/staff/:id", bindAndValidateReq[c.UserIdRequest](staffCon.GetStaff))
	e.PUT("/staff/:id", bindAndValidateReq[sc.UpdateStaffRequest](staffCon.UpdateStaff))
	e.PUT("/staff/:id/suspend", bindAndValidateReq[c.UserIdRequest](staffCon.SuspendStaff))
	e.PUT("/staff/:id/activate", bindAndValidateReq[c.UserIdRequest](staffCon.ActivateStaff))
	e.GET("/staffs", staffCon.GetAllStaffs)
}
