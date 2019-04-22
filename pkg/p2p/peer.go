package p2p

import (
	"fmt"
	"net"
)

type peer struct {
	id         PeerID
	connection net.Conn
}

func (p *peer) send(id PeerID, d []byte) error {
	fmt.Printf("Send in %d\n", p.id)
	fmt.Printf("Writing to socket of node %d: %s\n", id, d)
	return nil
}

func (p *peer) Write(d []byte) error {
	fmt.Printf("Writing to node %d: %s\n", p.id, d)
	return p.send(p.id, d)
}
