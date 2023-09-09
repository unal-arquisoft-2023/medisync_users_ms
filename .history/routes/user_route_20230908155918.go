package routes

import (
	"medysinc_user_ms/controllers"

	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo) {
	//pacientes

	e.POST("/patient", controllers.CreatePatient)
	e.GET("/patient/:userId", controllers.GetPatient)
	//doctors

	e.POST("/doctor", controllers.CreateDoctor)

	//admin y secretarios

	e.POST("/staff", controllers.CreateStaff)
	e.POST("/staff/:userId", controllers.GetStaffMember)
}
