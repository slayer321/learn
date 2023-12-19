package main

import (
	"fmt"
)

func main() {

	// server := NewServer()
	// server.Start()

	// go func() {
	// 	close(server.quitch)
	// }()

	// select {}
	loopTest()
}

func loopTest() {

	user := make(chan string, 10000)

	for i := 0; i < 10; i++ {
		go func(i int) {
			user <- fmt.Sprintf("user_%d", i)
		}(i)

	}

	defer close(user)

	for val := range user {
		fmt.Printf("value %v\n", val)
	}

}

type Server struct {
	users  map[string]string
	userch chan string
	quitch chan struct{}
}

func NewServer() *Server {
	return &Server{
		users:  make(map[string]string),
		userch: make(chan string),
		quitch: make(chan struct{}),
	}
}

func (s *Server) Start() {
	go s.loop()
}

func (s *Server) loop() {
	for {
		// user := <-s.userch
		// fmt.Printf("Printing user %s\n", user)
		// s.users[user] = user
		select {
		case msg := <-s.userch:
			fmt.Printf("Printing user %s\n", msg)
			s.users[msg] = msg
		case val := <-s.quitch:
			fmt.Printf("value of quitchannle is %+v\n", val)
			fmt.Println("quit the channel")
			return
		default:

		}
	}
}

func sendMessage(msg chan<- string) {
	msg <- "Hello"
}

func readMessage(msg <-chan string) {
	out := <-msg

	fmt.Println(out)
}
