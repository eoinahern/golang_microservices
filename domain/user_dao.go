package domain

import (
	"net/http"

	"github.com/eoinahern/golang_microservices/microservice_project/utils"
)

var dataMap = map[int64]*User{
	123: &User{ID: 123, FirstName: "eoin", LastName: "Ahern", Email: "hello@yahoo.com"},
}

//UserDao : instance of userDao that implements interface
var UserDao userDaoInterface

type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct {
}

//GetUser : fake getting user from DB
func (u *userDao) GetUser(id int64) (*User, *utils.ApplicationError) {

	if user := dataMap[id]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{Message: "user not found",
		StatusCode: http.StatusNotFound,
		Code:       "not_found"}
}
