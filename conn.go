package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type connection struct {
	// The websocket connection.
	ws *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
}

func (c *connection) reader() {
	for {
		var message *Message

		// TODO figure out why EOF makes everything explode
		err := c.ws.ReadJSON(&message)
		if message == nil {
			return
		}

		if err != nil {
			panic(err)
		}

		log.Println("Connection reader: ", message)
		message.route_place()

		res, err := json.Marshal(message)
		if err != nil {
			panic(err)
		}
		h.broadcast <- res
	}
	c.ws.Close()
}

func (c *connection) writer() {
	for message := range c.send {
		err := c.ws.WriteJSON(string(message))
		if err != nil {
			panic(err)
		}
	}
	c.ws.Close()
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(w, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		panic(err)
	}
	c := &connection{send: make(chan []byte, 256), ws: ws}
	h.register <- c
	defer func() { h.unregister <- c }()
	go c.writer()
	c.reader()
}
