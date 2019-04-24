package main

import (
	"p2p/pkg/p2p"
)

func main() {
	swarm := p2p.New(p2p.PeerID(0))
	swarm.Listen("0.0.0.0:2222")
}
