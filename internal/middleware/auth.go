package middleware

import (
	controllers "github.com/SoumyaCO/cr/internal/controller"
	"github.com/SoumyaCO/cr/internal/model"
)

// login middlware
// now returning boolean, but later it'll provide the jwt token
func Login(data model.SigninData) ([]string, error) {
	// query to the database, that this person exists or not
	userData, err := controllers.GetUserByEmail(data)
	if err != nil {
		return nil, err
	}

	if len(userData) > 0 {
		return userData, nil
	}

	return nil, err
}
