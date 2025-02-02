/*
Database access layer
logics related to database calls should be here
Related Files:
    - ./internal/middleware/auth.go
    - ./internal/controller/user.go
*/

package dal

import (
	"log"

	"github.com/SoumyaCO/cr/internal/db"
	"github.com/SoumyaCO/cr/internal/model"
)

type userDalInterface interface {
	CreateUser(model.User) (int64, error)
	UpdateUser(int, model.User) bool
	Deleteuser(int) bool
	GetUserByID(int) model.User
	GetUserByEmail(string) (model.User, error)
}

type UserDal struct {
	userDalInterface              // should implement userDalInterface
	DB               *db.Database // db for running queries
}


func (udal *UserDal) CreateUser(user model.User) (int64, error) {
	query := "INSERT INTO users (Email, Password) VALUES(?, ?)"
	result, queryError := udal.DB.Exec(query, user.Email, user.Password)
	if queryError != nil {
		log.Printf("CreateUserError:\t%v", queryError)
		return 0, queryError
	}
	rowsAffected, rowsAffectedError := result.RowsAffected()

	if rowsAffectedError != nil {
		log.Printf("CreateUserError:RowsAffectedError\t%v", rowsAffectedError)
		return 0, rowsAffectedError
	}
	return rowsAffected, nil
}

func (udal *UserDal) GetUserByEmail(email string) (model.User, error) {
	// create a model.User object
	user := model.User{}

	query := "SELECT * FROM users WHERE email=?"
	row := udal.DB.QueryRow(query, email)
    err := row.Scan(&user.Id, &user.Email, &user.Password)

	if err != nil {
		log.Printf("CreateUserError:RowsAffectedError\t%v", err)
		return user, err
	}

    return user, nil
}
