package controllers

import (
	"log"

	"github.com/SoumyaCO/cr/internal/dal"
	"github.com/SoumyaCO/cr/internal/db"
	"github.com/SoumyaCO/cr/internal/model"
)

type UserInterface interface {
	SignUp(model.User) (int64, error)
	DeleteUser(int) (int64, error)
}

type User struct {
	UserInterface
}

func (user User) CreateUser(userInfo model.User) (int64, error) {
    db, err := db.GetDatabaseConnection()
    if err != nil {
        log.Printf("DBConnectionError: %v", err)
        return 0, nil
    }
    udal := dal.UserDal{DB: db}
    rowsAffected, err := udal.CreateUser(userInfo)
    if err != nil {
        log.Printf("rowsAffectedError: %v", err)
        return 0, nil
    }
    return rowsAffected, nil
}
