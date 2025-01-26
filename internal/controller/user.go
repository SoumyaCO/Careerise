package controllers

import (
	"database/sql"
	"log"
	"os"

	"github.com/SoumyaCO/cr/internal/model"
	"github.com/go-sql-driver/mysql"
)

func CreateDB() (*sql.DB, error) {

    // TODO: load environment variable from an .env file

	config := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}

	// Get a database handle.
	var err error
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	return nil, nil
}

type Db *sql.DB

func GetUserByEmailAndPass(data model.SigninData) {
	err := Db.Query("SELECT password FROM guru WHERE email = ?", data.Email)

}
