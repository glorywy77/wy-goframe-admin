package model

import "time"

type UserLoginInput struct {
	UserName string
	Password string
}

type UserCreateInput struct {
	UserName    string
	Password    string
	Email       string
	Role        string
	Status      int
	Create_time time.Time
}

type UserCreateOutput struct {
  Result string
}