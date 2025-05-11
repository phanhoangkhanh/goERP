package test_repo_impl

import (
	"context"
	"myERP/banana"
	"myERP/db"
	"myERP/model"
	"time"
)

type TestRepoImpl struct {
	Sql *db.Sql
}

func (testRepo *TestRepoImpl) SaveUserTest(c context.Context, user model.User) (model.User, error) {
	statement := `
		INSERT INTO users(user_id, email, password, role, full_name, created_at, updated_at)
		VALUES(:user_id, :email, :password, :role, :full_name, :created_at, :updated_at)
	`
	user.CreateAt = time.Now()
	user.UpdateAt = time.Now()

	_, err := testRepo.Sql.Db.NamedExecContext(c, statement, user)

	if err != nil {
		return user, banana.TestSignUpFail
	}

	return user, nil

}
