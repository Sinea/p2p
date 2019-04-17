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
