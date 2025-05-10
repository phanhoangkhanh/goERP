package repo_impl

import (
	"context"
	"myERP/banana"
	"myERP/db"
	"myERP/model"
	"time"

	"github.com/lib/pq"
)

// work with DB so there is SQL
type UserRepoImpl struct {
	Sql *db.Sql
}

func (u *UserRepoImpl) SaveUser(context context.Context, user model.User) (model.User, error) {
	statement := ` INSERT INTO user(user_id, email, password, role, full_name, created_at, updated_at)
		VALUE(:user_id,:email, :password, :role, :full_name, :created_at, :updated_at )`
	user.CreateAt = time.Now()
	user.UpdateAt = time.Now()

	// NameExecContext parse user data into statement param
	_, err := u.Sql.Db.NamedExecContext(context, statement, user)
	if err != nil {
		// check error type is postgress error ? and this is expect unique of column data
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return user, banana.UserConflict
			}
		}
		//others errors , not pgError -> throw other message
		return user, banana.SignUpFail
	}
	// create database ok
	return user, nil

}
