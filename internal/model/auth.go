package model

// type SignupData struct {
// 	firstName  string
// 	lastName   string
// 	middleName string
// 	email      string
// 	gender     bool
// 	password   string
// }

// type ResetData struct {
// 	firstName string
// 	email     string
// }
//
// type auth interface {
// 	Signup(data SignupData)
// 	Signin(data SigninData)
// 	ResetPassword(data SigninData)
// }
//
// func SignUp(data SignupData) {
// }

// Sign in details
type SigninData struct {
	Email    string `json:"email"`    // email of the user
	Password string `json:"password"` // password of the user
}

