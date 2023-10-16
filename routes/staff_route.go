package routes

import (
	"medysinc_user_ms/controllers"

	"github.com/labstack/echo/v4"
)

func StaffRoute(e *echo.Echo) {

	//admin y secretarios

	e.POST("/staff", controllers.CreateStaff)
	e.GET("/staff/:staffId", controllers.GetStaffMember)
	e.PUT("/staff/:staffId", controllers.UpdateStaff)
}
