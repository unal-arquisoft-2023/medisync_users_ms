package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func bindAndValidateReq[T any](f func(c echo.Context, req T) error) echo.HandlerFunc {
	return func(c echo.Context) error {

		var req T
		if err := c.Bind(&req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := c.Validate(req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return f(c, req)
	}
}
