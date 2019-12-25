package domain

import (
	"net/http"

	"github.com/eoinahern/golang_microservices/microservice_project/utils"
)

var dataMap = map[int64]*User{
	123: &User{ID: 123, FirstName: "eoin", LastName: "Ahern", Email: "hello@yahoo.com"},
}

//GetUser : fake getting user from DB
func GetUser(id int64) (*User, *utils.ApplicationError) {

	if user := dataMap[id]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{Message: "user not found",
		StatusCode: http.StatusNotFound,
		Code:       ""}
}
