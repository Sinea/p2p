package p2p

import (
	"fmt"
)

type swarm struct {
	localID         PeerID
	nodes           map[PeerID]Node
	peers           map[PeerID]Peer
	peerRoutes      map[PeerID]Peer
	peerConnections map[PeerID][]PeerID
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

func (s *swarm) Node(id PeerID) Node {
	if p, ok := s.nodes[id]; ok {
		return p
	}
	return nil
}

func New(id PeerID) Swarm {
	return &swarm{
		localID:         id,
		nodes:           make(map[PeerID]Node),
		peerRoutes:      make(map[PeerID]Peer),
		peerConnections: make(map[PeerID][]PeerID),
	}
}

type Route struct {
	Distance uint32
	Through  PeerID
}
