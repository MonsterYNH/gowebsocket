package main

import (
	"github.com/gin-gonic/gin"
	"gowebsocket/controller"
)

func main() {
	engine := gin.Default()
	engine.GET("/", controller.WebSocket)
	message := controller.MessageController{}
	engine.POST("/message", message.SendMessage)
	engine.Run("0.0.0.0:3456")
}
