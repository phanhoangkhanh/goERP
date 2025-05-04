package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func WelcomeEveryOne(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, THIS IS KHANH'S NEW ERP PROJECT !")

}
