package model


type LoginCodeOutput struct {
	Code string
}




type UserLoginInput1 struct {
  Username string
  Password string	
	Code string
}


type UserLoginOutput struct {
	Token string
}
