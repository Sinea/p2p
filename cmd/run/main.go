package main

import "p2p/pkg/p2p"

func main() {
	swarm := p2p.New()
	swarm.Node(0).Write([]byte("hello"))
}
