package app

import (
	"github.com/eoinahern/golang_microservices/microservice_project/controllers"
)

func mapURLs() {
	router.GET("/users/:user_id", controllers.GetUser)
}
