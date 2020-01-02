package services

import (
	"net/http"
	"testing"

	"github.com/eoinahern/golang_microservices/microservice_project/domain"
	"github.com/eoinahern/golang_microservices/microservice_project/utils"
	"github.com/stretchr/testify/assert"
)

var UserDaoMock userDaoMock

type userDaoMock struct {
}

func (m *userDaoMock) GetUser(id int64) (*domain.User, *utils.ApplicationError) {
	return nil, nil
}

func TestGetUserNotFound(t *testing.T) {

	user, err := UserService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
}

func TestGetNoErr(t *testing.T) {

	user, err := UserService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
}
