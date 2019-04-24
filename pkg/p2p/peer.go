package p2p

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
)

type peer struct {
	id         PeerID
	connection net.Conn
	buffer     []byte
}

func (p *peer) read() error {
	buffer := make([]byte, 1000)
	if n, err := p.connection.Read(buffer); err != nil {
		return err
	} else {
		p.buffer = append(p.buffer, buffer[:n]...)
	}

	if len(p.buffer) < 13 {
		return nil
	}

	if p.buffer[0] != header {
		return errors.New("invalid message header")
	}

	id := binary.BigEndian.Uint64(p.buffer[1:9])
	size := binary.BigEndian.Uint32(p.buffer[9:13])

	fmt.Printf("Received %d bytes from %d\n", size, id)
	if t := p.buffer[13:]; uint32(len(t)) >= size {
		fmt.Println(string(p.buffer[13 : 13+size]))
	}

	return nil
}

func (p *peer) send(id PeerID, d []byte) error {
	fmt.Printf("Send in %d\n", p.id)
	fmt.Printf("Writing to socket of node %d: %s\n", id, d)

	tmp := make([]byte, 13+len(d))
	tmp[0] = header
	binary.BigEndian.PutUint64(tmp[1:9], uint64(id))
	binary.BigEndian.PutUint32(tmp[9:13], uint32(len(d)))
	copy(tmp[13:], d)
	if _, err := p.connection.Write(tmp); err != nil {
		return err
	}
	return nil
}

func (p *peer) Write(d []byte) error {
	fmt.Printf("Writing to node %d: %s\n", p.id, d)
	return p.send(p.id, d)
}
