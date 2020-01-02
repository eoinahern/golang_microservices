package services

import (
	"net/http"
	"testing"

	"github.com/eoinahern/golang_microservices/microservice_project/domain"
	"github.com/eoinahern/golang_microservices/microservice_project/utils"
	"github.com/stretchr/testify/assert"
)

var getUserfunction func(int64) (*domain.User, *utils.ApplicationError)

type userDaoMock struct {
}

func (m *userDaoMock) GetUser(id int64) (*domain.User, *utils.ApplicationError) {
	return getUserfunction(id)
}

func init() {
	domain.UserDao = &userDaoMock{}
}

func TestGetUserNotFound(t *testing.T) {

	getUserfunction = func(id int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
	}

	user, err := UserService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
}

func TestGetNoErr(t *testing.T) {

	getUserfunction = func(id int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{ID: 1}, nil
	}
	user, err := UserService.GetUser(1)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, user.ID)
}
