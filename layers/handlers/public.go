package handlers

import "github.com/labstack/echo/v4"

func Public(e *echo.Echo) {
	e.Static("/", "public")
}
