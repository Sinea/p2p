package main

import (
	"p2p/pkg/p2p"
	"time"
)

func main() {
	swarm := p2p.New(p2p.PeerID(time.Now().UnixNano()))
	swarm.Listen("0.0.0.0:2222")
}
