package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {

	user, err := GetUser(0)

	assert.Nil(t, user, "we are not expecting a user!!")
	assert.NotNil(t, err, "error should not be nil??")
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "user not found", err.Message)

	if user != nil {
		t.Error("not expecting user id 0")
	}

	if err == nil {
		t.Error("we were expecting an error")
	}

	if err.StatusCode != http.StatusNotFound {
		t.Error("expecting status not found?")
	}
}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)

	assert.NotNil(t, user, "user should be returned")
	assert.Nil(t, err)
	assert.EqualValues(t, "eoin", user.FirstName)
	assert.EqualValues(t, "Ahern", user.LastName)
	assert.EqualValues(t, "hello@yahoo.com", user.Email)
}
