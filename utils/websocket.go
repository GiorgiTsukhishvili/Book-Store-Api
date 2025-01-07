package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/models"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/scripts"
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
)

func HandleWebSocket(ctx *gin.Context) {
	conn, err := Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		log.Println("Error upgrading connection:", err)
		return
	}

	defer conn.Close()

	claims := scripts.GetUserClaims(ctx)

	UserID := strconv.FormatUint(uint64(claims.UserID), 10)

	user := &User{
		ID:   UserID,
		Conn: conn,
	}
	UsersMu.Lock()
	Users[UserID] = user
	UsersMu.Unlock()

	log.Printf("User connected: %s\n", UserID)

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Printf("User disconnected: %s\n", UserID)
			UsersMu.Lock()
			delete(Users, UserID)
			UsersMu.Unlock()
			break
		}
	}
}

func SendMessage(UserID string, message models.Notification) error {
	UsersMu.RLock()
	User, exists := Users[UserID]
	UsersMu.RUnlock()

	if !exists {
		return fmt.Errorf("User with ID %s not found", UserID)
	}

	data, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("error marshalling message: %v", err)
	}

	User.Conn.WriteMessage(websocket.TextMessage, data)

	log.Printf("Message sent to %s \n", UserID)
	return nil
}
