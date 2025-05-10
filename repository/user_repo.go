package repository

import (
	"context"
	"myERP/model"
)

type UserRepo interface {
	SaveUser(context context.Context, user model.User) (model.User, error)
}
