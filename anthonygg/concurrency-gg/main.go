package main

import (
	"fmt"
	"time"
)

type Message struct {
	From    string
	Payload string
}

type Server struct {
	msgch  chan Message
	quitch chan struct{}
}

func (s *Server) ListenAndServer() {

running:
	for {
		select {

		case msg := <-s.msgch:
			fmt.Printf("Message is from %s and payload is %s\n", msg.From, msg.Payload)
		case <-s.quitch:
			fmt.Printf("quitting the channel\n")
			break running
		default:
		}
	}
}

func sendMessageToServer(msgch chan Message, payload string) {
	msg := Message{
		From:    "Sachin Maurya",
		Payload: payload,
	}

	msgch <- msg

}

func graceFullQuit(quitch chan struct{}) {
	close(quitch)
}

func main() {

	server := &Server{
		msgch:  make(chan Message),
		quitch: make(chan struct{}),
	}

	go server.ListenAndServer()
	go func() {
		time.Sleep(2 * time.Second)
		sendMessageToServer(server.msgch, "THis is important message")
	}()

	go func() {
		time.Sleep(4 * time.Second)
		graceFullQuit(server.quitch)
	}()

	select {}
}
