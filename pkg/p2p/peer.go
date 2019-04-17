package p2p

import "fmt"

type peer struct {
	id PeerID
}

func (p *peer) send(id PeerID, d []byte) error {
	fmt.Printf("Writing to socket of node %d: %s\n", id, d)
	return nil
}

func (p *peer) Write(d []byte) error {
	fmt.Printf("Writing to node %d: %s\n", p.id, d)
	return p.send(p.id, d)
}
