package main

import (
	"fmt"
	"p2p/pkg/p2p"
)

type application struct {
	network p2p.Swarm
}

func (a *application) HandleMessage(d []byte) {
	fmt.Printf(">> %s\n", string(d))
}

func (a *application) Connected() {
	fmt.Printf("I'm now connected\n")
}

func (a *application) Disconnected() {
	fmt.Printf("I disconected\n")
}

func (a *application) Joined(id p2p.NodeID) {
	fmt.Printf("Node with id %d joined\n", id)
	node, err := a.network.Node(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := node.Write([]byte("Welcome")); err != nil {
		fmt.Printf("Error writing to node: %s\n", err)
	}
}

func (*application) Left(id p2p.NodeID) {
	fmt.Printf("Node with id %d left\n", id)
}
