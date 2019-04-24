package p2p

const (
	header = 0xAC
)

type PeerID uint64

type Node interface {
	Write([]byte) error
}

type Peer interface {
	send(PeerID, []byte) error
	read() error
}

type Router interface {
	Route(PeerID, []byte) error
}

type Swarm interface {
	Node(PeerID) (Node, error)
	Listen(address string) error
	Join(address string) error

	setConnections(PeerID, []PeerID)
}
