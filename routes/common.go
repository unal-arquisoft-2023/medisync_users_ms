package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func bindReq[T any](f func(c echo.Context, req T) error) echo.HandlerFunc {
	return func(c echo.Context) error {

		var req T
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return f(c, req)
	}
}
