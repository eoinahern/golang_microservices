package services

import "github.com/eoinahern/golang_microservices/microservice_project/domain"
import "github.com/eoinahern/golang_microservices/microservice_project/utils"

//GetUser : gets a user
func GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userID)
}
