package engine

import (
	"golang.org/x/net/websocket"
)

// Client model
type Client struct {
	ws   *websocket.Conn
	name string
}

// NewClient is a constructor
func NewClient(ws *websocket.Conn, name string) *Client {
	return &Client{ws, name}
}
