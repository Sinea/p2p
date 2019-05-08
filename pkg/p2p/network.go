package p2p

import (
	"errors"
	"fmt"
	"net"
	"p2p/pkg/p2p/protocol"
)

type network struct {
	token           []byte
	application     Application
	localID         NodeID
	nodes           map[NodeID]Node
	peers           map[NodeID]Peer
	peerRoutes      map[NodeID]Peer
	peerConnections map[NodeID][]NodeID
}

func (n *network) Join(address string) error {
	connection, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}
	allocatedID, allocatorID, err := joinNetwork(connection, n.token)
	if err != nil {
		return err
	}
	n.localID = NodeID(allocatedID)
	n.peers[NodeID(allocatorID)] = &peer{
		localID:  n.localID,
		protocol: protocol.New(connection, 1024, header),
		handler:  n.application,
	}
	fmt.Printf("Accepted by %d\n", allocatorID)
	fmt.Printf("Allocated id %d\n", allocatedID)
	n.application.Connected()
	return nil
}

func (n *network) Listen(address string) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	// For now
	go n.read()

	for {
		connection, err := listener.Accept()
		if err != nil {
			return err
		}
		allocatedID, err := acceptNode(connection, n.token)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Accepted new node")
			peer := &peer{
				id:       NodeID(allocatedID),
				localID:  n.localID,
				protocol: protocol.New(connection, 1024, header),
				handler:  n.application,
			}
			n.peers[NodeID(allocatedID)] = peer
			n.nodes[NodeID(allocatedID)] = peer
			n.application.Joined(NodeID(allocatedID))
		}
	}
}

func (n *network) read() {
	for {
		for id, p := range n.peers {
			if err := p.read(); err != nil {
				fmt.Printf("error reading from peer %d\n", id)
				delete(n.peers, id)
				n.application.Left(p.ID())
			}
		}
	}
}

func (n *network) setConnections(from NodeID, to []NodeID) {
	n.peerConnections[from] = to
	n.updateRoutingTable()
}

func (n *network) Route(id NodeID, b []byte) error {
	fmt.Printf("Routing message to %d\n", id)
	if p, ok := n.peerRoutes[id]; ok {
		return p.send(id, message, b)
	}
	return fmt.Errorf("no route found to peer %d", id)
}

func (n *network) Node(id NodeID) (Node, error) {
	if p, ok := n.nodes[id]; ok {
		return p, nil
	}
	return nil, errors.New("not found")
}

func New(id NodeID, app Application, token []byte) Network {
	return &network{
		token:           token,
		application:     app,
		localID:         id,
		nodes:           make(map[NodeID]Node),
		peerRoutes:      make(map[NodeID]Peer),
		peerConnections: make(map[NodeID][]NodeID),
		peers:           make(map[NodeID]Peer),
	}
}
