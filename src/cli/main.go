package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

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

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("My name is: ")
	var name string
	if scanner.Scan() {
		name = scanner.Text()
	}

	fmt.Println("Type ':q' for exit")

	client := engine.NewClient(ws, name)
	go client.Listen()

	var text string
	for {
		if scanner.Scan() {
			text = scanner.Text()
			if text == ":q" {
				ws.Close()
				break
			} else if text != "" {
				client.Say(text)
			}
		}
	}
}
