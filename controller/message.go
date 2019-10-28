package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MessageController struct {}

type MessageData struct {
	UserId string `json:"user_id"`
	Type int `json:"type"`
	Num int `json:"num"`
}

func (controller *MessageController) SendMessage(c *gin.Context) {
	var messageData MessageData
	if err := c.BindJSON(&messageData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	if len(messageData.UserId) == 0 || messageData.Type == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errors.New("ERROR: parameter error"),
		})
		return
	}
	client, exist := Manage.GetClient(messageData.UserId)
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "client is out of line",
		})
		return
	}
	if err := client.SendMessage(messageData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
