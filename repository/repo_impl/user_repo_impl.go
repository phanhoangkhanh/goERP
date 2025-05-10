package repo_impl

import (
	"context"
	"myERP/db"
	"myERP/model"
	"time"
)

// work with DB so there is SQL
type UserRepoImpl struct {
	sql *db.Sql
}

func (u *UserRepoImpl) SaveUser(context context.Context, user model.User) (model.User, error) {
	statement := ` INSERT INTO user(user_id, email, password, role, full_name, created_at, updated_at)
		VALUE(:user_id,:email, :password, :role, :full_name, :created_at, :updated_at )`
	user.CreateAt = time.Now()
	user.UpdateAt = time.Now()

	// NameExecContext parse user data into statement param
	u.sql.Db.NamedExecContext(context, statement, user)

}
