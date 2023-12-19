package main

import (
	"fmt"
	"log"

	"github.com/slayer321/casstorage/p2p"
)

func OnPeer(peer p2p.Peer) error {
	//fmt.Println("doing some login with the peer outside of TCP transport")
	peer.Close()
	return nil
}

func main() {
	opts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
		OnPeer:        OnPeer,
	}
	tr := p2p.NewTCPTransport(opts)

	go func() {
		for {
			msg := <-tr.Consume()
			fmt.Printf("%+v\n", msg)
		}
	}()

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	//time.Sleep(time.Second * 20)

	select {}
}
