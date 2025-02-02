package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Database struct {
	*sql.DB
}

var instance *Database
var once sync.Once

func GetDatabaseConnection() (*Database, error) {
	once.Do(func() {
		var db *sql.DB

        // important: relative path of .env file
		godotenv.Load("../../.env")

        // print environment variables
		fmt.Println("DB User\t", os.Getenv("LOCAL_DB_USER"))
		fmt.Println("DB Pass\t", os.Getenv("LOCAL_DB_PASS"))
		fmt.Println("DB Name\t", os.Getenv("LOCAL_DB_NAME"))

		cfg := mysql.Config{
			User:   os.Getenv("LOCAL_DB_USER"),
			Passwd: os.Getenv("LOCAL_DB_PASS"),
			Net:    "tcp",
			Addr:   os.Getenv("LOCAL_DB_PORT"),
			DBName: os.Getenv("LOCAL_DB_NAME"),
		}

		var err error
		db, err = sql.Open("mysql", cfg.FormatDSN())
		if err != nil {
			log.Fatalf("Error in DB connection \n %v", err)
		}
		instance = &Database{db}
	})
	return instance, nil
}
