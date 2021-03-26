package notification

import (
	"flag"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
)

var addr = flag.String("addr", "localhost:8080", "http service address")


func SendNotification(message string) {

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	err = c.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		log.Println("write:", err)
		return
	}
}
