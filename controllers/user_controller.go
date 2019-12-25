package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/eoinahern/golang_microservices/microservice_project/services"
	"github.com/eoinahern/golang_microservices/microservice_project/utils"
)

//GetUser : returns a user
func GetUser(resp http.ResponseWriter, req *http.Request) {

	userID, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)

	if err != nil {

		userErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad request",
		}

		marshaledErr, _ := json.Marshal(userErr)
		resp.WriteHeader(userErr.StatusCode)
		resp.Write([]byte(marshaledErr))
		return
	}

	myuser, userErr := services.GetUser(userID)

	if userErr != nil {
		jsonErr, _ := json.Marshal(userErr)
		resp.WriteHeader(userErr.StatusCode)
		resp.Write([]byte(jsonErr))
		return
	}

	jsonUser, _ := json.Marshal(myuser)
	resp.WriteHeader(http.StatusOK)
	resp.Header().Add("Content-Type", "application/json")
	resp.Write(jsonUser)

}
