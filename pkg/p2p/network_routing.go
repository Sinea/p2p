package p2p

import "math"

// Ugly code. Refactor

func (n *network) setConnections(from NodeID, to []NodeID) {
	n.peerConnections[from] = to
	n.updateRoutingTable()
}

func (n *network) updateRoutingTable() {
	n.peerRoutes = make(map[NodeID]Peer)
	for id := range n.nodes {
		if id == n.localID || n.peers[id] != nil {
			continue
		}
		routes := n.findRoute(n.localID, id)
		for _, x := range routes {
			if len(x) > 0 {
				n.peerRoutes[id] = n.nodes[x[1]].(Peer)
				break
			}
		}
	}
}

func (n *network) findRoute(from, to NodeID) [][]NodeID {
	l := n.paths(from, to, []NodeID{from})
	routes := make([][]NodeID, 0)
	size := math.MaxInt32
	for _, t := range l {
		tLen := len(t)
		if tLen < size {
			size = len(t)
			routes = make([][]NodeID, 0)
		} else if tLen > size {
			continue
		}
		routes = append(routes, t)
	}

	return routes
}

func (n *network) paths(from, to NodeID, visited []NodeID) [][]NodeID {
	peers := n.peerConnections[from]
	result := make([][]NodeID, 0)
	collected := make([][]NodeID, 0)
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
		z := n.paths(v, to, t)
		for _, zz := range z {
			collected = append(collected, zz)
		}
	}
	return collected
}

func contains(needle NodeID, haystack []NodeID) bool {
	for _, t := range haystack {
		if t == needle {
			return true
		}
	}
	return false
}
