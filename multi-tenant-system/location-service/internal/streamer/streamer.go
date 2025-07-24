package streamer

import (
	"encoding/json"
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

var conn *websocket.Conn
var connected bool

// LocationPayload defines the structure of location data sent
type LocationPayload struct {
	TenantID  string  `json:"tenant_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Connect establishes the WebSocket connection once
func Connect() {
	u := url.URL{Scheme: "ws", Host: "streaming-service:8083", Path: "/ws"}
	var err error
	conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Printf("WebSocket connection failed: %v", err)
		connected = false
		return
	}
	log.Println("Connected to streaming service via WebSocket")
	connected = true
}

// SendLocation pushes the data to streaming-service
func SendLocation(payload LocationPayload) {
	if conn == nil {
		log.Println("No active WebSocket connection")
		return
	}

	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal payload: %v", err)
		return
	}

	maxRetries := 3
	for i := range maxRetries {
		if !connected || conn == nil {
			log.Println("WebSocket disconnected. Attempting to reconnect...")
			Connect()
			if !connected {
				log.Printf("Retry %d/%d: Unable to reconnect", i+1, maxRetries)
				time.Sleep(2 * time.Second)
				continue
			}
		}

		err := conn.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			log.Printf("WebSocket write failed: %v", err)
			connected = false
			time.Sleep(2 * time.Second)
			continue
		}

		return
	}

	log.Println("Failed to send location after retries. Giving up.")
}
