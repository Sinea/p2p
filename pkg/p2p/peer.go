package p2p

import (
	"p2p/pkg/p2p/protocol"
)

type peer struct {
	id       NodeID
	handler  MessageHandler
	localID  NodeID
	protocol *protocol.Protocol
}

func (p *peer) ID() NodeID {
	return p.id
}

func (p *peer) send(id NodeID, command uint8, data []byte) error {
	return p.protocol.Write(command, packData(id, data))
}

func (p *peer) read() (command uint8, payload []byte, err error) {
	return p.protocol.Read()
}

func (p *peer) Write(data []byte) error {
	return p.send(p.id, message, data)
}
