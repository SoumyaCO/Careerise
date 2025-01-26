package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/SoumyaCO/cr/internal/model"
	// "github.com/SoumyaCO/cr/internal/middleware"
)


func LoginHandler(rw http.ResponseWriter, r *http.Request) {
	// 1. get the data from frontend
	// 2. login() function from middleware
	// 3. write back the cookie into the cookie-header

	decoder := json.NewDecoder(r.Body)
	var data model.SigninData
	err := decoder.Decode(&data)
	if err != nil {
		log.Printf("[DATA READING ERROR]\tError while reading the data: %v", err)
	}

    
}
