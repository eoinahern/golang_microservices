package app

import (
	"fmt"
	"net/http"

	"github.com/eoinahern/golang_microservices/microservice_project/controllers"
)

//StartApp : starts the application
func StartApp() {

	fmt.Println("app starting ......")
	http.HandleFunc("/users", controllers.GetUser)
	http.ListenAndServe(":9090", nil)

}
