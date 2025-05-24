package controllers

import (
	"github.com/goyalaadi12071998/app/models"
	"github.com/goyalaadi12071998/app/services"
)

var UserController usercontroller

type usercontroller struct {
	userService services.IUserService
}

func InitialiseUserController(userService services.IUserService) {
	UserController = usercontroller{
		userService: userService,
	}
}

func (u usercontroller) CreateUser(name string, email string, password string) *models.User {
	user := u.userService.CreateUser(name, email, password)
	return user
}
