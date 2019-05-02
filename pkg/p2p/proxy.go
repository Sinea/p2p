package p2p

type proxy struct {
	id     NodeID
	router Router
}

func (p *proxy) Write(d []byte) error {
	//fmt.Printf("Will route message\n")
	return p.router.Route(p.id, d)
}
