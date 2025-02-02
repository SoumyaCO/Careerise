package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/SoumyaCO/cr/internal/middleware"
	"github.com/SoumyaCO/cr/internal/model"
)

func LoginHandler(rw http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var data model.LoginInfo
	err := decoder.Decode(&data)
	if err != nil {
		log.Printf("JsonDecodingError:\t%v", err)
	}
	status, err := middleware.LoginMiddleware(data.Email, data.Password)
	if err != nil {
		log.Printf("LoginMiddlewareError:\t%v", err)
	}
	if status {
		rw.Write([]byte("Authorized!"))
	} else {
		rw.Write([]byte("UnAuthorized!!!"))
	}
	return
}
