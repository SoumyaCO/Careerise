package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Database struct {
	careerise *sql.DB
}

type dbConfigEnv struct {
	user   string
	pass   string
	dbname string
}

func NewDatabaseConnection() (*sql.DB, error) {
	godotenv.Load()
	var envConfig dbConfigEnv

	// Based on the environmen strings will be different
	if os.Getenv("ENVIRONMENT") == "LOCAL" {
		envConfig = dbConfigEnv{
			user:   os.Getenv("LOCAL_DB_USER"),
			pass:   os.Getenv("LOCAL_DB_PASS"),
			dbname: os.Getenv("LOCAL_DB_NAME"),
		}
	} else {
		envConfig = dbConfigEnv{
			user:   os.Getenv("PROD_DB_USER"),
			pass:   os.Getenv("PROD_DB_PASS"),
			dbname: os.Getenv("PROD_DB_NAME"),
		}
	}

	config := mysql.Config{
		User:   envConfig.user,
		Passwd: envConfig.pass,
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: envConfig.dbname,
	}

	var err error
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatalf("Error in DB connection \n %v", err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatalf("Ping Error \n%v", pingErr)
	}
	return db, err
}
