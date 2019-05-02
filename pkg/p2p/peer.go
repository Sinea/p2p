package p2p

import (
	"encoding/binary"
	"fmt"
	"net"
)

const (
	bufferSize = 1000
)

type peer struct {
	id         NodeID
	connection net.Conn
	buffer     []byte
	handler    MessageHandler
	localID    NodeID
	protocol   *Proto
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

func (p *peer) send(id NodeID, command uint8, d []byte) error {
	//fmt.Printf("Send in %d\n", p.id)
	//fmt.Printf("Writing to socket of node %d: %s\n", id, d)

	tmp := make([]byte, 8+len(d))
	tmp[0] = header
	tmp[1] = command
	// Destination node
	binary.BigEndian.PutUint16(tmp[2:4], uint16(id))
	// Payload length
	binary.BigEndian.PutUint32(tmp[4:8], uint32(len(d)))
	// Write payload
	copy(tmp[8:], d)

	if _, err := p.connection.Write(tmp); err != nil {
		return err
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
