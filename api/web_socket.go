package api

import (
	db "Chat-system_golang/db/sqlc"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			// Allow connections from any origin
			return true
		},
	}

	// UserRegistry stores the connected users
	UserRegistry = struct {
		sync.RWMutex
		users map[int32]*User
	}{
		users: make(map[int32]*User),
	}
)

// User represents a user connected to the WebSocket server
type User struct {
	ID   int32
	conn *websocket.Conn
}

type AckMsg struct {
	Msg         string
	IsDelivered bool
}

type MessageToSend struct {
	ID         int32     `json:"id"`
	Message    string    `json:"message"`
	SenderID   int32     `json:"sender_id"`
	ReceiverID int32     `json:"receiver_id"`
	SentAt     time.Time `json:"sent_at"`
}

// HandleWebSocket handles the WebSocket connections
func (server *Server) handleWebSocket(ctx *gin.Context) {
	// Upgrade HTTP connection to WebSocket
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println("Error upgrading connection to WebSocket:", err)
		return
	}

	log.Println("User Successfully Connected...")

	query := ctx.Request.URL.Query()

	senderId := query.Get("sender_id")

	userID, _ := strconv.Atoi(senderId)
	newUserID := int32(userID)

	// Create a new user
	user := &User{
		ID:   newUserID,
		conn: conn,
	}

	// Register the user
	UserRegistry.Lock()
	UserRegistry.users[newUserID] = user
	UserRegistry.Unlock()

	// Handle incoming messages
	for {
		// Read message from the WebSocket connection
		_, msgBytes, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		// Parse the message
		var message db.Message
		if err := json.Unmarshal(msgBytes, &message); err != nil {
			log.Println("Error parsing message:", err)
			continue
		}

		message.SenderID = newUserID

		// Handle the message
		SendMessageToUser(server, conn, ctx, message)
	}

	// Unregister the user when the connection is closed
	UserRegistry.Lock()
	delete(UserRegistry.users, newUserID)
	UserRegistry.Unlock()
}

// SendMessageToUser sends a message to a specific user
func SendMessageToUser(server *Server, webConn *websocket.Conn, ctx *gin.Context, message db.Message) {
	UserRegistry.RLock()
	defer UserRegistry.RUnlock()

	dbMessage, err := server.store.StoreMessage(ctx, db.StoreMessageParams{
		Message:    message.Message,
		SenderID:   message.SenderID,
		ReceiverID: message.ReceiverID,
	})
	if err != nil {
		log.Println("Cannot send message")
		return
	}

	recipient, ok := UserRegistry.users[message.ReceiverID]
	if !ok {

		ackByte, err := json.Marshal(AckMsg{
			Msg:         dbMessage.Message,
			IsDelivered: false,
		})
		if err != nil {
			log.Println("Error sending message to user:", err)
			return
		}

		if err := UserRegistry.users[message.SenderID].conn.WriteMessage(websocket.TextMessage, ackByte); err != nil {
			log.Println("Error sending message to user:", err)
		}

		log.Println("Recipient not found:", message.ReceiverID)
		return
	}

	// updating successful delivery status of message that is sent
	_, _ = server.store.UpdateMessageDelivery(ctx, dbMessage.ID)

	msgBytes, err := json.Marshal(MessageToSend{
		ID:         dbMessage.ID,
		Message:    dbMessage.Message,
		SenderID:   dbMessage.SenderID,
		ReceiverID: dbMessage.ReceiverID,
		SentAt:     dbMessage.SentAt,
	})
	if err != nil {
		log.Println("Error sending message to user:", err)
		return
	}

	if err := recipient.conn.WriteMessage(websocket.TextMessage, msgBytes); err != nil {
		log.Println("Error sending message to receiver:", err)
	}

	time.Sleep(2 * time.Second)

	ackByte, err := json.Marshal(AckMsg{
		Msg:         dbMessage.Message,
		IsDelivered: true,
	})
	if err != nil {
		log.Println(err)
		return
	}

	if err := UserRegistry.users[message.SenderID].conn.WriteMessage(websocket.TextMessage, ackByte); err != nil {
		log.Println("Error sending acknowledgement message to sender:", err)
	}
}
