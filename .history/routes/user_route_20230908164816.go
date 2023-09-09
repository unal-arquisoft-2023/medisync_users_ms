package routes

import (
	"medysinc_user_ms/controllers"

	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo) {
	//pacientes

	e.POST("/patient", controllers.CreatePatient)
	e.GET("/patient/:patientId", controllers.GetPatient)
	e.PUT("/patient/:patientId", controllers.UpdatePatient)
	//doctors

	e.POST("/doctor", controllers.CreateDoctor)
	e.GET("/doctor/:doctorId", controllers.GetDoctor)

	//admin y secretarios

	e.POST("/staff", controllers.CreateStaff)
	e.GET("/staff/:userId", controllers.GetStaffMember)
	e.PUT("/staff/:userId", controllers.UpdateStaffMember)
}
