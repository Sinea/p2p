package p2p

import (
	"encoding/binary"
	"p2p/pkg/p2p/protocol"
)

const (
	bufferSize = 1000
)

type peer struct {
	id       NodeID
	handler  MessageHandler
	localID  NodeID
	protocol *protocol.Protocol
}

func (p *peer) send(id NodeID, command uint8, data []byte) error {
	return p.protocol.Write(command, packData(id, data))
}

func (p *peer) ID() NodeID {
	return p.id
}

func (p *peer) read() (command uint8, payload []byte, err error) {
	return p.protocol.Read()
}

func (p *peer) Write(data []byte) error {
	return p.send(p.id, message, data)
}

func packData(id NodeID, data []byte) []byte {
	message := make([]byte, 2+len(data))
	binary.BigEndian.PutUint16(message[:2], uint16(id))
	copy(message[2:], data)
	return message
}

func unpackData(data []byte) (NodeID, []byte) {
	id := NodeID(binary.BigEndian.Uint16(data[:2]))
	return id, data[2:]
}
