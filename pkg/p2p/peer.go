package p2p

import (
	"encoding/binary"
	"fmt"
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

func (p *peer) ID() NodeID {
	return p.id
}

func (p *peer) read() error {
	cmd, data, err := p.protocol.Read()

	if err != nil {
		return err
	}

	switch cmd {
	case message:
		id, body := unpackData(data)
		if id == p.localID {
			p.handler.HandleMessage(body)
		} else {
			fmt.Println("Just pass")
		}
	}

	return nil
}

func (p *peer) Write(d []byte) error {
	//fmt.Printf("Writing to node %d: %s\n", p.id, d)
	m := packData(p.id, d)
	return p.protocol.Write(message, m)
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
