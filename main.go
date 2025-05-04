package main

import (
	"myERP/db"
	"myERP/handler"

	"github.com/labstack/echo/v4"
)

func main() {

	//Connect Pg
	sql := &db.Sql{

		Host:     "172.21.137.52",
		Port:     5432,
		UserName: "khanhadmin",
		Password: "123",
		DbName:   "goerp",
	}
	sql.Connect()
	defer sql.Close()

	e := echo.New()

	e.GET("/", handler.WelcomeEveryOne)

	e.GET("/user/sign-in", handler.HandlerSignIn)
	e.GET("/user/testing-response-json", handler.HandlerTestResponseObj)

	e.Logger.Fatal(e.Start(":1323"))
}
