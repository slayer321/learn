package main

import (
	"fmt"
	"log"
	"net"
)

type message struct {
	quit chan struct{}
	con  string
}

type Server struct {
	listenAddr string
	ln         net.Listener
	quitch     message
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
		quitch:     message{quit: make(chan struct{})},
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}
	fmt.Println("listening on", s.listenAddr)
	defer ln.Close()

	s.ln = ln
	go s.acceptLoop()
	//fmt.Println("Before the send channel")
	<-s.quitch.quit
	//fmt.Println("After the send channel")
	return nil
}

func (s *Server) acceptLoop() {
	//fmt.Println("Inside the acceptLoop")
	for {
		//fmt.Println("Inside the acceptLoop for loop")
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Printf("accept error: %v\n", err)
			continue
		}
		fmt.Printf("accept connection from %v \n", conn.RemoteAddr().String())
		s.quitch.con = conn.RemoteAddr().String()
		go s.readLoop(conn)

	}
}

func (s *Server) readLoop(conn net.Conn) {
	//fmt.Println("Inside the readLoop")
	defer conn.Close()
	buf := make([]byte, 2048)
	for {
		//fmt.Println("Inside the readLoop for loop")
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("read error: %v\n", err)
			continue
		}

		msg := buf[:n]
		//fmt.Println(string(msg))
		fmt.Printf("connection (%v):  %v\n", s.quitch.con, string(msg))
	}
}

func main() {
	server := NewServer(":12345")
	log.Fatal(server.Start())
}
