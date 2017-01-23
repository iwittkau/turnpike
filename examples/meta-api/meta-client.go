package main

import (
	"log"

	"github.com/iwittkau/turnpike"
)

func main() {
	c, err := turnpike.NewWebsocketClient(turnpike.JSON, "ws://localhost:8000/", nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected to router")
	_, err = c.JoinRealm("turnpike.examples", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("joined realm")
	c.ReceiveDone = make(chan bool)

	onJoin := func(args []interface{}, kwargs map[string]interface{}) {
		log.Println("session joined:", args[0])
	}
	if err := c.Subscribe("wamp.session.on_join", nil, onJoin); err != nil {
		log.Fatalln("Error subscribing to channel:", err)
	}

	onLeave := func(args []interface{}, kwargs map[string]interface{}) {
		log.Println("session left:", args[0])
	}
	if err := c.Subscribe("wamp.session.on_leave", nil, onLeave); err != nil {
		log.Fatalln("Error subscribing to channel:", err)
	}

	log.Println("listening for meta events")
	<-c.ReceiveDone
	log.Println("disconnected")
}
