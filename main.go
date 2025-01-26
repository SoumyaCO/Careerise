package main

import (
	"log"
	"net/http"
	"os"

	"github.com/SoumyaCO/cr/internal/db"
	"github.com/joho/godotenv"
)

func main() {
	mux := http.NewServeMux()

    // Database connection testing
    _, dberror := db.NewDatabaseConnection()
    if dberror != nil {
        log.Println("DB ERROR")
    }

	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("<h1>Home Route</h1>"))
	})
	mux.HandleFunc("/profile", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("<h1>Profile Route</h1>"))
	})

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env variables \n%v", err)
	}

	port := os.Getenv("PORT")

	log.Printf("Server started at port: %v \n", port)
	err = http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatalf("Error starting the server \n%v", err)
	}
}
