package app

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

//StartApp : starts the application
func StartApp() {

	mapURLs()
	router.Run(":9090")

}

func init() {
	router = gin.Default()
}
