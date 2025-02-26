package handler

import "github.com/labstack/echo/v4"

func Health(e echo.Context) error {
	return e.String(200, "healthy")
}
