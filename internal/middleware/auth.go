package middleware

import (
	"log"

	"github.com/SoumyaCO/cr/internal/dal"
	"github.com/SoumyaCO/cr/internal/db"
)

// just get the password, providing the email and udal and match it
// implement tests.. after that ..
// ..build the middleware, then the encryption part
func LoginMiddleware(email, pass string) (bool, error) {
	// for now it'll look unncecessary. but integrated with server logics and encryptions
	// it'll come to a stage to show it's true purpose
    db, DBerror := db.GetDatabaseConnection()
    if DBerror != nil {
        log.Printf("DBConnectionError: %v", DBerror)
    }
    udal := dal.UserDal{
        DB: db,
    }
    user, err := udal.GetUserByEmail(email)
	if err != nil {
		log.Printf("Auth:GettingUserError;\t %v", err)
		return false, err
	}

    // check
    if user.Password == pass {
        return true, nil
    }
	return false, nil
}
