package routers

import (
	"github.com/francescoforesti/go-demo/gin/handlers"
	"github.com/francescoforesti/go-demo/gin/logging"
	constants "github.com/francescoforesti/go-demo/gin/utils"
	"github.com/gin-gonic/gin"
	"os"
)

func InitializeServerPort() string {
	var serverPort = os.Getenv(constants.SERVER_PORT_ENV_VAR)
	if len(serverPort) == 0 {
		serverPort = constants.SERVER_PORT
	}
	return serverPort
}

func InitializeRoutes(router *gin.Engine) {

	router.RedirectTrailingSlash = true
	router.RedirectFixedPath = true

	api := router.Group("/api/v1/")

	logging.Debug("Creating routes '/api/v1'...")

	api.POST("/strings", handlers.PostString)
	api.GET("/strings/reversed", handlers.GetReversedString)
}
