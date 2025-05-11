package handler

import (
	"myERP/model"
	"myERP/model/req"
	"myERP/repository"
	"myERP/security"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type TestHandler struct {
	TestRepoImp repository.TestRepo
}

func (testQuery *TestHandler) TestHandlerSignUp(c echo.Context) error {
	req := req.ReqSignUp{}

	// 1.Bind req with model
	if err := c.Bind(&req); err != nil {
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

	userID, err := uuid.NewUUID()
	hash := security.HashAndSalt([]byte(req.Password))
	role := model.SUPERADMIN.StringRole()

	if err != nil {
		return c.JSON(http.StatusForbidden, model.Response{
			StatusCode: http.StatusForbidden,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	user := model.User{
		UserId:   userID.String(),
		FullName: req.FullName,
		Email:    req.Email,
		Password: hash,
		Role:     role,
		Token:    "",
	}

	//Execute Repo
	user, err = testQuery.TestRepoImp.SaveUserTest(c.Request().Context(), user)

	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    "Có lỗi gì đó",
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Done- thành công rổi nhé",
		Data:       user,
	})

}
