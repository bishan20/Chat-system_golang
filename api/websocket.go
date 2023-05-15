package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func reader(conn *websocket.Conn) {

	for {
		// reading messasge from Client
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(string(msg))

		// writing message back to Client
		if err := conn.WriteMessage(messageType, msg); err != nil {
			log.Println(err)
			return
		}
	}
}

func webSocketHandler(w http.ResponseWriter, r *http.Request) {

	// allowing any connections into this web socket endpoint regardless of what the origin of that connection is
	wsupgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// upgrading the HTTP connection to web-socket connection
	wsConn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %+v", err)
		return
	}

	defer wsConn.Close()

	log.Println("Client Successfully Connected....")

	reader(wsConn)

}
