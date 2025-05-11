package repo_impl

import (
	"context"
	"fmt"
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

	statement := `
		INSERT INTO users(user_id, email, password, role, full_name, created_at, updated_at)
		VALUES(:user_id, :email, :password, :role, :full_name, :created_at, :updated_at)
	`
	user.CreateAt = time.Now()
	user.UpdateAt = time.Now()
	fmt.Println("USER??", user)

	// NameExecContext parse user data into statement param
	_, err := u.Sql.Db.NamedExecContext(context, statement, user)
	if err != nil {
		// check error type is postgress error ? and this is expect unique of column data
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return user, banana.UserConflict
			}
		}
		fmt.Println("VÀO DEN DAY KO??")
		//others errors , not pgError -> throw other message
		return user, banana.SignUpFail
	}
	// create database ok
	return user, nil

}
