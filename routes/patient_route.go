package routes

import (
	"medysinc_user_ms/controllers/patient"
	"net/http"

	"github.com/labstack/echo/v4"
)

func bind[T any](f func(c echo.Context, req T) error) echo.HandlerFunc {
	return func(c echo.Context) error {

		var req T
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return f(c, req)
	}
}

func PatientRoute(e *echo.Echo, patCon *patient.PatientController) {
	//pacientes
	e.POST("/patient", bind[patient.CreatePatientRequest](patCon.CreatePatient))
}
