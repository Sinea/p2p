package main

import (
	"fmt"
	"log"
	"p2p/pkg/p2p"
)

type app struct {
	network p2p.Swarm
}

func (*app) HandleMessage(d []byte) {
	fmt.Printf(">> %s\n", string(d))
}

func (*app) Connected() {
	fmt.Printf("I'm now connected\n")
}

func (*app) Disconnected() {
	fmt.Printf("I disconnected\n")
}

func (a *app) Joined(id p2p.NodeID) {
	fmt.Printf("Node with id %d joined\n", id)
	if node, err := a.network.Node(id); err == nil {
		node.Write([]byte("Welcome"))
	} else {
		fmt.Println(err)
	}
}

func (*app) Left(id p2p.NodeID) {
	fmt.Printf("Node with id %d left\n", id)
}

func main() {
	app := &app{}
	app.network = p2p.New(p2p.NodeID(0), app)
	if err := app.network.Listen("0.0.0.0:2222"); err != nil {
		log.Fatalf("Error listening: %s\n", err)
	}
}
