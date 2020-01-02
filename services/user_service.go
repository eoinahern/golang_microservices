package services

import "github.com/eoinahern/golang_microservices/microservice_project/domain"
import "github.com/eoinahern/golang_microservices/microservice_project/utils"

//UserService : userservice struct
type userService struct {
}

var UserService userService = userService{}

//GetUser : gets a user
func (u *userService) GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return domain.UserDao.GetUser(userID)
}
