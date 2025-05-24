package models

import (
	"time"
)

var Users map[*User]bool = map[*User]bool{}

type User struct {
	Name      string
	Email     string
	Password  string
	CreatedAt int64 //Unix Timestamp
	UpdatedAt int64 //Unix Timestamo
}

type IUser interface {
	SaveUserToDB()
}

func (u *User) SaveUserToDB() *User {
	Users[u] = true
	return u
}

func NewUser(name string, email string, password string) *User {
	return &User{
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
}
