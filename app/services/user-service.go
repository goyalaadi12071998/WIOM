package services

import "github.com/goyalaadi12071998/app/models"

var userService UserService

type UserService struct {
	fakeDbConnection string
}

func InitialiseUserService() IUserService {
	userService = UserService{}
	return userService
}

type IUserService interface {
	CreateUser(name string, email string, password string) *models.User
}

func (u UserService) CreateUser(name string, email string, password string) *models.User {
	user := models.NewUser(name, email, password)
	user.SaveUserToDB()
	return user
}
