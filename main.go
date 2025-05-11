package main

import (
	"myERP/db"
	"myERP/handler"
	"myERP/repository/repo_impl"
	"myERP/repository/test_repo_impl"

	"myERP/router"

	"github.com/labstack/echo/v4"
)

func main() {

	//CONNECT DATABASE - PG
	sql := &db.Sql{

		Host:     "172.21.137.52",
		Port:     5432,
		UserName: "khanhadmin",
		Password: "123",
		DbName:   "goerp",
	}
	sql.Connect()
	defer sql.Close()

	//CONNECT DATABASE TWO - PG
	sqlTwo := &db.Sql{
		Host:     "172.21.137.52",
		Port:     5432,
		UserName: "khanhadmin",
		Password: "123",
		DbName:   "goerp2",
	}
	sqlTwo.Connect()
	defer sqlTwo.Close()

	//ROUTING PART
	e := echo.New()
	router.GenerateAPICallHandler(e)
	// call router with API object-struct with 2 param inject obj-struct ( echo + userHandler)
	userRepoImplement := repo_impl.UserRepoImpl{Sql: sql}
	userHandler := handler.UserHandler{
		UserRepo: &userRepoImplement,
	}
	testHandler := handler.TestHandler{
		TestRepoImp: &test_repo_impl.TestRepoImpl{Sql: sql},
	}
	api := router.API{
		Echo:        e,
		UserHandler: userHandler,
		TestHandler: testHandler,
	}
	api.SetupRouter()

	// e.GET("/", handler.WelcomeEveryOne)

	e.Logger.Fatal(e.Start(":1323"))

	//THIS IS FROM TESTBRANCH
}
