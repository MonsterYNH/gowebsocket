package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gowebsocket/model"
	"log"
	"net/http"
)

var Manage = model.CreateManage()

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebSocket(c *gin.Context) {
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("ERRPR: upgrader websocker failed, error:", err)
		return
	}
	defer ws.Close()


	client := model.CreateClient(ws)
	Manage.AddClient("123", client)
	client.ListenClient()
}
