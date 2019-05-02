package p2p

const (
	header   uint8 = 0xAC
	message  uint8 = 1
	join     uint8 = 2
	accepted uint8 = 3
	rejected uint8 = 4
)

type NodeID uint16

type Node interface {
	Write([]byte) error
}

type Peer interface {
	ID() NodeID
	send(NodeID, uint8, []byte) error
	read() error
}

type Router interface {
	Route(NodeID, []byte) error
}

type Swarm interface {
	Node(NodeID) (Node, error)
	Listen(address string) error
	Join(address string) error

	setConnections(NodeID, []NodeID)
}

type MessageHandler interface {
	HandleMessage([]byte)
}

type Application interface {
	MessageHandler

	// Self status
	Connected()
	Disconnected()

	// Other nodes
	Joined(NodeID)
	Left(NodeID)
}
