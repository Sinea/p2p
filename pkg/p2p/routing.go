package p2p

import "math"

// Ugly code. Refactor

func (s *swarm) updateRoutingTable() {
	s.peerRoutes = make(map[PeerID]Peer)
	for id := range s.nodes {
		if id == s.localID || s.peers[id] != nil {
			continue
		}
		routes := s.findRoute(s.localID, id)
		for _, x := range routes {
			if len(x) > 0 {
				s.peerRoutes[id] = s.nodes[x[1]].(Peer)
				break
			}
		}
	}
}

func (s *swarm) findRoute(from, to PeerID) [][]PeerID {
	l := s.paths(from, to, []PeerID{from})
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

func (s *swarm) paths(from, to PeerID, visited []PeerID) [][]PeerID {
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
		z := s.paths(v, to, t)
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
