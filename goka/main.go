package main

import (
	"github.com/francescoforesti/go-demo/goka/handlers"
	"github.com/francescoforesti/go-demo/goka/logging"
	"github.com/francescoforesti/go-demo/goka/routers"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func main() {
	logging.InitializeLoggers()
	logging.Debug("this is debug")
	logging.Info("this is info")
	logging.Warn("this is warn")
	logging.Error("this is error")

	gin.SetMode(gin.DebugMode)
	router = gin.Default()
	routers.InitializeRoutes(router)
	handlers.InitializeHandlers()
	router.Run(routers.InitializeServerPort())

}
