package p2p

import "fmt"

type proxy struct {
	id     PeerID
	router Router
}

func (p *proxy) Write(d []byte) error {
	fmt.Printf("Will route message\n")
	return p.router.Route(p.id, d)
}
