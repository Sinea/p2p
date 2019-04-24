package main

import (
	"fmt"
	"p2p/pkg/p2p"
	"time"
)

func main() {
	swarm := p2p.New(p2p.PeerID(1))
	go swarm.Listen("0.0.0.0:1111")
	swarm.Join("0.0.0.0:2222")
	if node, err := swarm.Node(0); err == nil {
		node.Write([]byte("hello world!"))
	} else {
		fmt.Println(err)
	}
	time.Sleep(time.Hour)
}
