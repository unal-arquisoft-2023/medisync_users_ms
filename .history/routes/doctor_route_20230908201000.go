package routes

import (
	"medysinc_user_ms/controllers"

	"github.com/labstack/echo/v4"
)

func StaffRoute(e *echo.Echo) {

	//doctors

	e.POST("/doctor", controllers.CreateDoctor)
	e.GET("/doctor/:doctorId", controllers.GetDoctor)
	e.PUT("/doctor/:doctorId", controllers.UpdateDoctor)
	e.GET("/doctors", controllers.GetAllDoctors)
	e.GET("/doctors", controllers.GetDoctorsBySpecialty)

}
