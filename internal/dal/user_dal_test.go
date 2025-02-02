package dal

import (
	"fmt"
	"testing"

	"github.com/SoumyaCO/cr/internal/db"
	"github.com/SoumyaCO/cr/internal/model"
)

func TestCreateUser(t *testing.T) {
	db, err := db.GetDatabaseConnection()
	fmt.Println("DBOBJECT: ", db)
	if err != nil {
		t.Fatalf("ErrorGettingDBConnection: %v", err)
	}
	udal := UserDal{DB: db}
	user := model.User{
		Email:    "soumyadip@duck.com",
		Password: "pass",
	}
	rows, err := udal.CreateUser(user)
	if err != nil {
		t.Fatalf("ErrorCreatingUser: %v", err)
	}

	if int(rows) == 0 {
		t.Fatal("ZeroRowAffectedError: zero rows affected in database.")
	}
}
