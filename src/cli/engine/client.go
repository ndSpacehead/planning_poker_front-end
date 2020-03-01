package engine

import (
	"fmt"
	"io"
	"log"

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

// Say some text
func (c *Client) Say(text string) {
	c.sendCh <- text
}

// Listen "write" and "read" events
func (c *Client) Listen() {
	go c.listenWrite()
	c.listenRead()
}

func (c *Client) listenWrite() {
	for {
		select {
		case text := <-c.sendCh:
			websocket.JSON.Send(c.ws, &Message{c.name, text})
		}
	}
}

func (c *Client) listenRead() {
	for {
		select {
		default:
			var msg Message
			err := websocket.JSON.Receive(c.ws, &msg)
			if err == io.EOF {
				log.Panicln("Empty message")
			} else if err != nil {
				log.Panicln(err)
			} else {
				if msg.Author != c.name {
					fmt.Println(msg.String())
				}
			}
		}
	}
}
