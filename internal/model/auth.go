package model

type SignupInfo struct {
    Name string `json:"name"`
    Email string `json:"email"`
    Password string `json:"password"`
}


type LoginInfo struct {
    Email string `json:"email"`
    Password string `json:"password"`
}


