package p2p

import (
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
	s := New(0)
	n := s.Node(1)

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
		localID:         0,
		peerRoutes:      make(map[PeerID]Peer),
		peerConnections: make(map[PeerID][]PeerID),
	}

	s.nodes = map[PeerID]Node{
		1: &peer{1},
		2: &peer{2},
		3: &proxy{3, s},
	}

	s.peers = map[PeerID]Peer{
		1: &peer{1},
		2: &peer{2},
	}

	s.setConnections(0, []PeerID{1, 2})
	s.setConnections(1, []PeerID{0, 2, 3})
	s.setConnections(2, []PeerID{0, 1})
	s.setConnections(3, []PeerID{1})

	s.Node(3).Write([]byte("hello world"))
}
