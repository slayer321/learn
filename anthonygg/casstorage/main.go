package main

import (
	"fmt"
	"log"

	"github.com/slayer321/casstorage/p2p"
)

func main() {
	fmt.Println("Hello world")
	start := p2p.NewTCPTransport(":12345")
	if err := start.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	//time.Sleep(time.Second * 20)

	select {}
}
