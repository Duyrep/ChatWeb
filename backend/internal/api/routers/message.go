package routers

import (
	"DuyrepWebsiteBackend/internal/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getMessages(c *gin.Context) {
	amountStr := c.Query("amount")
	if amountStr == "" {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	amount, err := strconv.ParseInt(amountStr, 10, 64)
	if amount > 100 {
		c.JSON(http.StatusOK, gin.H{"Error": "Limited to 100 messages"})
	} else if err != nil {
		c.JSON(http.StatusOK, gin.H{"Error": "Invalid Parameter"})
	} else {
		c.JSON(http.StatusOK, database.GetMessages("global", amount))
	}
}

func sendMessage(c *gin.Context) {
	var message database.Message
	if err := c.ShouldBindJSON(&message); err != nil {
		return
	}
	database.SendMessage("global", message)
}
