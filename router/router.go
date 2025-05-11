package router

import (
	"myERP/handler"

	"github.com/labstack/echo/v4"
)

type API struct {
	Echo        *echo.Echo
	UserHandler handler.UserHandler
	TestHandler handler.TestHandler
}

func (api *API) SetupRouter() {
	api.Echo.GET("/user/sign-in", api.UserHandler.HandlerSignIn)
	api.Echo.POST("/user/sign-up", api.UserHandler.HandlerSignUp)
	api.Echo.POST("/test/sign-up", api.TestHandler.TestHandlerSignUp)

	api.Echo.GET("/user/testing-response-json", api.UserHandler.HandlerTestResponseObj)
}

// For testing
func GenerateAPICallHandler(e *echo.Echo) *echo.Echo {

	e.GET("/", handler.WelcomeEveryOne)
	return e
}
