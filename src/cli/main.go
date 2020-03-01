package main

import (
	"fmt"
	"log"

	"golang.org/x/net/websocket"

	"./engine"
)

var origin = "http://localhost/"
var url = "ws://localhost:8080/entry"

func main() {
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Print("My name is: ")
	var name string
	fmt.Scanln(&name)

	client := engine.NewClient(ws, name)

	ws.Close()
}
