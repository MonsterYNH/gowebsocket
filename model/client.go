package model

import (
	"github.com/gorilla/websocket"
	"log"
	"strings"
	"time"
)

type Client struct {
	conn       *websocket.Conn
	updateTime int64
}

type WebSocketResponse struct {
	Ping string      `json:"ping"`
	Data interface{} `json:"data,omitempty"`
}

func CreateClient(conn *websocket.Conn) *Client {
	client := &Client{
		conn:       conn,
		updateTime: time.Now().Unix(),
	}
	return client
}

func (client *Client) ListenClient() {
	defer client.conn.Close()

	go client.heartBreak()
	for {
		_, message, err := client.conn.ReadMessage()
		if err != nil {
			if strings.Contains(err.Error(), "use of closed network connection") {
				break
			} else {
				log.Println("ERROR: read message failed, error: ", err)
				break
			}
		}
		if string(message) == "ping" {
			message = []byte("pong")
		} else {
			log.Println("ERROR: wrong option with client")
			break
		}
		client.updateTime = time.Now().Unix()
		if err := client.SendMessage(nil); err != nil {
			log.Println("ERROR: write message failed, error: ", err)
			break
		}
	}
}

func (client *Client) heartBreak() {
	timer := time.NewTicker(time.Second * 10)
	for {
		<-timer.C
		if time.Now().Add(-time.Second * 5).Unix() > client.updateTime {
			log.Println("客户端响应超时,断开")
			client.conn.Close()
			break
		}
	}
}

func (client *Client) SendMessage(data interface{}) error {
	return client.conn.WriteJSON(WebSocketResponse{
		Ping: "pong",
		Data: data,
	})
}
