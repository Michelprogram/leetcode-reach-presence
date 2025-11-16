package handlers

import (
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

type Event = string

const (
	UPDATE_TITLE Event = "UpdateTitle"
	UPDATE_TIMER Event = "UpdateTimer"
)

func (wh WebsocketHandler) Controller(w http.ResponseWriter, r *http.Request) {

	up := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			origin := r.Header.Get("Origin")
			if origin == "" {
				return true
			}
			if strings.HasPrefix(origin, "chrome-extension://") {
				return true
			}
			if strings.HasPrefix(origin, "http://localhost:") || strings.HasPrefix(origin, "https://localhost:") {
				return true
			}
			return false
		},
	}

	c, err := up.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				slog.Info("WebSocket closed by client", "err", err)
			} else if websocket.IsUnexpectedCloseError(err, websocket.CloseAbnormalClosure, websocket.CloseGoingAway) {
				slog.Warn("WebSocket unexpected close", "err", err)
			} else {
				slog.Warn("WebSocket read error", "err", err)
			}
			break
		}

		var msg Message

		err = json.Unmarshal(message, &msg)
		if err != nil {
			c.WriteMessage(websocket.TextMessage, []byte("Wrong data structure"))
			continue
		}

		log.Printf("Received message: %+v\n", msg)

		wh.Queue <- msg

	}
}
