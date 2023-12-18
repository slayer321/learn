package main

import (
	"log"

	"github.com/slayer321/casstorage/p2p"
)

func main() {
	opts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
	}
	start := p2p.NewTCPTransport(opts)
	if err := start.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	//time.Sleep(time.Second * 20)

	select {}
}
