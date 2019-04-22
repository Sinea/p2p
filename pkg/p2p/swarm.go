package p2p

import (
	"errors"
	"fmt"
	"net"
)

type swarm struct {
	localID         PeerID
	nodes           map[PeerID]Node
	peers           map[PeerID]Peer
	peerRoutes      map[PeerID]Peer
	peerConnections map[PeerID][]PeerID
}

func (s *swarm) Join(address string) error {
	connection, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}
	s.handshake(connection)
	return nil
}

func (s *swarm) Listen(address string) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	for {
		connection, err := listener.Accept()
		if err != nil {
			return err
		}
		go s.handshake(connection)
	}
}

func (s *swarm) setConnections(from PeerID, to []PeerID) {
	s.peerConnections[from] = to
	s.updateRoutingTable()
}

func (s *swarm) Route(id PeerID, b []byte) error {
	fmt.Printf("Routing message to %d\n", id)
	if p, ok := s.peerRoutes[id]; ok {
		return p.send(id, b)
	}
	return fmt.Errorf("no route found to peer %d", id)
}

func (s *swarm) Node(id PeerID) (Node, error) {
	if p, ok := s.nodes[id]; ok {
		return p, nil
	}
	return nil, errors.New("not found")
}

func New(id PeerID) Swarm {
	return &swarm{
		localID:         id,
		nodes:           make(map[PeerID]Node),
		peerRoutes:      make(map[PeerID]Peer),
		peerConnections: make(map[PeerID][]PeerID),
	}
}
