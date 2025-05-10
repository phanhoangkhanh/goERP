package handler

import (
	"myERP/model"
	"myERP/model/req"
	"myERP/security"
	"net/http"

	validator "github.com/go-playground/validator/v10"
	uuid "github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
}

func (u *UserHandler) HandlerSignUp(c echo.Context) error {
	// request from FE with type of ReqSigUp struct
	req := req.ReqSignUp{}

	//1. check Bind Parse Request + Bind request from FE to variable req
	if err := c.Bind(&req); err != nil {
		//if Bind parse fail -> return Json to FE with struct response
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	//2. Check Validator Request
	validate := validator.New()

	if err := validate.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	hash := security.HashAndSalt([]byte(req.Password))
	role := model.MEMBER.StringRole()
	userID, err := uuid.NewUUID()
	if err != nil {
		return c.JSON(http.StatusForbidden, model.Response{
			StatusCode: http.StatusForbidden,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	return nil
}

func (u *UserHandler) HandlerSignIn(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func (u *UserHandler) HandlerTestResponseObj(c echo.Context) error {

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
