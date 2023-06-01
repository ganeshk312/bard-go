package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ganeshk312/bard-go/bard"

	"github.com/gin-gonic/gin"
)

type Message struct {
	SessionID string `json:"session_id"`
	Message   string `json:"message"`
}

var initBool bool
var chatbot *bard.Chatbot

func ask(c *gin.Context) {
	log.Println("Someone asked me something")
	log.Println(c)
	var message Message
	if err := c.BindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	log.Println(message)

	sessionID := message.SessionID
	userAuthKey := os.Getenv("USER_AUTH_KEY")

	// Check if the user has defined an auth key,
	// If so, check if the auth key in the header matches it.
	if userAuthKey != "" && c.GetHeader("Authorization") != userAuthKey {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization key"})
		return
	}

	// Execute your code without authenticating the resource
	if !initBool {
		chatbot, _ = bard.NewChatbot(sessionID)
		initBool = true
	}
	fmt.Printf("chatbot %+v", chatbot)
	response, err := chatbot.Ask(message.Message)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, response)
}

func main() {
	fmt.Println("Entered server")
	r := gin.Default()

	r.POST("/ask", ask)
	if err := r.Run(":8000"); err != nil {
		fmt.Println("Failed to start server:", err)
	} else {
		fmt.Println("Listening at 8080")
	}
}
