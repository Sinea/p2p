package p2p

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSwarm_Node(t *testing.T) {
	s := &swarm{}
	s.nodes = map[PeerID]Node{
		0: &peer{0},
		1: &proxy{1, s},
		2: &proxy{2, s},
	}
	s.peerRoutes = map[PeerID]Peer{
		1: s.nodes[0].(Peer),
	}

	e1 := s.Node(0).Write([]byte("message for node 0"))
	e2 := s.Node(1).Write([]byte("message for node 1"))

	assert.NoError(t, e1)
	assert.NoError(t, e2)
}

func TestSwarm_Node2(t *testing.T) {
	s := New()
	n := s.Node(0)

	assert.Nil(t, n)
}

func TestSwarm_InvalidRoute(t *testing.T) {
	s := &swarm{}
	s.nodes = map[PeerID]Node{
		0: &peer{0},
		1: &proxy{1, s},
	}
	err := s.Node(1).Write([]byte("message"))

	assert.Error(t, err)
}

func TestSwarm_BuildRoutingTable(t *testing.T) {
	s := &swarm{
		nodes: map[PeerID]Node{
			0: nil,
			1: nil,
			2: nil,
			3: nil,
			4: nil,
		},
		peerRoutes:      make(map[PeerID]Peer),
		peerConnections: make(map[PeerID][]PeerID),
	}

	s.setConnections(0, []PeerID{1, 5})
	s.setConnections(1, []PeerID{0, 2, 5})
	s.setConnections(2, []PeerID{1, 3})
	s.setConnections(3, []PeerID{2, 4, 5})
	s.setConnections(4, []PeerID{3, 6})
	s.setConnections(5, []PeerID{0, 3, 1, 6})
	s.setConnections(6, []PeerID{4, 5})

	fmt.Printf("Route from %d to %d goes through %d\n", 0, 2, s.findRoute(PeerID(6), PeerID(1)))
	//for i := 0; i < 6; i++ {
	//	for j := 0; j < 6; j++ {
	//		if i == j {
	//			continue
	//		}
	//		fmt.Printf("Route from %d to %d goes through %d\n", i, j, s.findRoute(PeerID(i), PeerID(j)))
	//	}
	//}

}
