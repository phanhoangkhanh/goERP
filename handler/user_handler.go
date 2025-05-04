package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandlerSignIn(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func HandlerTestResponseObj(c echo.Context) error {

	type User struct {
		Email string `json:"myEmail"`
		Name  string `json:"pureName"`
		Age   int    `json:"realAge"`
	}
	userTest := User{
		Email: "phkhanh1188@gmail.com",
		Name:  "phanhoangkhanh",
		Age:   34,
	}
	return c.JSON(http.StatusOK, userTest)
}
