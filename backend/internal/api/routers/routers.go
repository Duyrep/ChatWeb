package routers

import (
	"DuyrepWebsiteBackend/internal/api/middleware"
	"DuyrepWebsiteBackend/internal/api/websocket"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func Run() {
	addr := os.Getenv("API_ADDRESS")
	router = gin.Default()
	router.Use(cors.New(middleware.GetCorsConfig()))
	setupRouter()
	router.Run(addr)
}

func setupRouter() {
	router.GET("/ws", websocket.WebSocket)

	router.GET("/ping", ping)
	router.GET("/get_messages", getMessages)

	router.POST("/send_message", sendMessage)
}
