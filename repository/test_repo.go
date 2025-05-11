package repository

import (
	"context"
	"myERP/model"
)

type TestRepo interface {
	SaveUserTest(c context.Context, user model.User) (model.User, error)
}
