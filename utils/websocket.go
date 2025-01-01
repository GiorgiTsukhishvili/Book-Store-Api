package utils

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type User struct {
	ID   string
	Conn *websocket.Conn
}

var (
	Users    = make(map[string]*User)
	UsersMu  sync.RWMutex
	Upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return strings.Contains(os.Getenv("ALLOWED_ORIGINS"), r.Header.Get("Origin"))
		},
	}
	MessageQueue = make(chan struct {
		ClientID string
		Message  string
	})
)

func HandleWebSocket(c *gin.Context) {
	conn, err := Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Error upgrading connection:", err)
		return
	}
	defer conn.Close()

	userID := c.Query("id")
	if userID == "" {
		log.Println("User ID is required")
		return
	}

	user := &User{
		ID:   userID,
		Conn: conn,
	}
	UsersMu.Lock()
	Users[userID] = user
	UsersMu.Unlock()

	log.Printf("User connected: %s\n", userID)

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Printf("User disconnected: %s\n", userID)
			UsersMu.Lock()
			delete(Users, userID)
			UsersMu.Unlock()
			break
		}
	}
}

func SendMessage(UserID string, message string) error {
	UsersMu.RLock()
	User, exists := Users[UserID]
	UsersMu.RUnlock()

	if !exists {
		return fmt.Errorf("User with ID %s not found", UserID)
	}

	err := User.Conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		log.Printf("Error sending message to %s: %v\n", UserID, err)
		return err
	}
	log.Printf("Message sent to %s: %s\n", UserID, message)
	return nil
}
