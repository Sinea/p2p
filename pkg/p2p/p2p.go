package p2p

const (
	// Control header
	header uint8 = 0xAC
	// Commands
	message  uint8 = 1
	join     uint8 = 2
	accepted uint8 = 3
	rejected uint8 = 4
	// Error messages
	errorJoiningNetwork   = "error joining network"
	invalidJoinToken      = "invalid join token"
	noJoinMessageReceived = "no join message received"
)

type NodeID uint16

type Node interface {
	Write([]byte) error
}

type Peer interface {
	ID() NodeID
	read() (command uint8, payload []byte, err error)
	send(NodeID, uint8, []byte) error
}

type router interface {
	Route(NodeID, []byte) error
}

type Network interface {
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
