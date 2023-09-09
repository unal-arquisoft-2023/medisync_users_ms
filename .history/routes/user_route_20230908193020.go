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
	e.GET("/patients", controllers.GetAllPatients)
	//doctors

	e.POST("/doctor", controllers.CreateDoctor)
	e.GET("/doctor/:doctorId", controllers.GetDoctor)
	e.PUT("/doctor/:doctorId", controllers.UpdateDoctor)
	e.GET("/doctors", controllers.GetAllDoctors)

	//admin y secretarios

	e.POST("/staff", controllers.CreateStaff)
	e.GET("/staff/:userId", controllers.GetStaffMember)
	e.PUT("/staff/:userId", controllers.UpdateStaffMember)
}
