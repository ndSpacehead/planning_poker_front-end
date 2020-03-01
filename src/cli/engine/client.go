package engine

import (
	"golang.org/x/net/websocket"
)

// Client model
type Client struct {
	ws        *websocket.Conn
	name      string
	sendCh    chan string
	receiveCh chan *Message
}

// NewClient is a constructor
func NewClient(ws *websocket.Conn, name string) *Client {
	sendCh := make(chan string)
	receiveCh := make(chan *Message)

	return &Client{
		ws,
		name,
		sendCh,
		receiveCh,
	}
}
