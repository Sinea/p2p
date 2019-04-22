package p2p

import (
	"fmt"
	"math"
)

type swarm struct {
	localID         PeerID
	nodes           map[PeerID]Node
	peerRoutes      map[PeerID]Peer
	peerConnections map[PeerID][]PeerID
}

func (s *swarm) setConnections(from PeerID, to []PeerID) {
	s.peerConnections[from] = to
	s.updateRoutingTable()
}

func (s *swarm) updateRoutingTable() {
	//for id, _ := range s.nodes {
	//	fmt.Printf("Route to %d goes through %d\n", id, s.findRoute(s.localID, id))
	//}
}

func (s *swarm) findRoute(from, to PeerID) [][]PeerID {
	l := s.lengths(from, to, []PeerID{from})
	routes := make([][]PeerID, 0)
	size := math.MaxInt32
	for _, t := range l {
		tLen := len(t)
		if tLen < size {
			size = len(t)
			routes = make([][]PeerID, 0)
		} else if tLen > size {
			continue
		}
		routes = append(routes, t)
	}

	return routes
}

func (s *swarm) lengths(from, to PeerID, visited []PeerID) [][]PeerID {
	peers := s.peerConnections[from]
	result := make([][]PeerID, 0)
	collected := make([][]PeerID, 0)
	for _, t := range peers {
		if t == to {
			collected = append(collected, append(visited, to))
			continue
		}
		if contains(t, visited) {
			continue
		}
		result = append(result, append(visited, t))
	}
	for _, t := range result {
		v := t[len(t)-1]
		z := s.lengths(v, to, t)
		for _, zz := range z {
			collected = append(collected, zz)
		}
	}
	return collected
}

func contains(needle PeerID, haystack []PeerID) bool {
	for _, t := range haystack {
		if t == needle {
			return true
		}
	}
	return false
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

func New() Swarm {
	return &swarm{
		nodes:           make(map[PeerID]Node),
		peerRoutes:      make(map[PeerID]Peer),
		peerConnections: make(map[PeerID][]PeerID),
	}
}

type Route struct {
	Distance uint32
	Through  PeerID
}
