package p2p

type proxy struct {
	id     NodeID
	router router
}

func (p *proxy) Write(d []byte) error {
	return p.router.Route(p.id, d)
}
