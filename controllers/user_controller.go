package controllers

import (
	"net/http"
	"strconv"

	"github.com/eoinahern/golang_microservices/microservice_project/services"
	"github.com/eoinahern/golang_microservices/microservice_project/utils"
	"github.com/gin-gonic/gin"
)

//GetUser : returns a user
func GetUser(cont *gin.Context) {

	userID, err := strconv.ParseInt(cont.Param("user_id"), 10, 64)

	if err != nil {

		userErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad request",
		}

		cont.JSON(http.StatusBadRequest, userErr)
		return
	}

	myuser, userErr := services.UserService.GetUser(userID)

	if userErr != nil {
		cont.JSON(userErr.StatusCode, userErr)
		return
	}

	cont.JSON(http.StatusOK, myuser)

}
