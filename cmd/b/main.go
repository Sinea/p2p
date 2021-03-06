package main

import (
	"fmt"
	"log"
	"p2p/pkg/p2p"
	"time"
)

func main() {
	app := &application{}
	app.network = p2p.New(p2p.NodeID(99), app, []byte("token123"))

	go func() {
		if err := app.network.Listen("0.0.0.0:1111"); err != nil {
			log.Fatalf("Error listening: %s\n", err)
		}
	}()

	if err := app.network.Join("0.0.0.0:2222"); err != nil {
		fmt.Printf("Error joining: %s\n", err)
	} else {
		time.Sleep(time.Hour)
	}
}
