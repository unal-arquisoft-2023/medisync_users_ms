package main

import (
	"medysinc_user_ms/configs"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	//run database
	configs.ConnectDB()

	e.Logger.Fatal(e.Start(":6000"))
}
