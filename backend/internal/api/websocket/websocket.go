package websocket

import (
	"DuyrepWebsiteBackend/internal/database"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Websocket error:", err)
	}
	go handle(conn)
}

func handle(conn *websocket.Conn) {
	cs := database.WatchChannel("global")

	defer conn.Close()
	defer cs.Close(context.TODO())

	for cs.Next(context.TODO()) {
		var event bson.M
		if err := cs.Decode(&event); err != nil {
			break
		}

		err := conn.WriteMessage(websocket.TextMessage, []byte("Changed"))
		if err != nil {
			break
		}
	}
}
