package protocol

import (
	"encoding/binary"
	"errors"
	"net"
)

// 1b CONTROL 0xAC	1
// 1b COMMAND		2
// 4b LENGTH		2:6
// Nb PAYLOAD		6:

type Protocol struct {
	connection     net.Conn
	buffer         []byte
	header         uint8
	readBufferSize uint
}

func (p *Protocol) Write(command uint8, body []byte) (err error) {
	message := make([]byte, 6)
	message[0] = p.header
	message[1] = command

	if body != nil {
		binary.BigEndian.PutUint32(message[2:6], uint32(len(body)))
		message = append(message, body...)
	}

	n, err := p.connection.Write(message)

	if err != nil {
		return
	}

	if n != len(message) {
		err = errors.New("wrote incomplete message")
	}

	return
}

func (p *Protocol) Read() (command uint8, payload []byte, err error) {
	buffer := make([]byte, p.readBufferSize)
	n, err := p.connection.Read(buffer)

	if err != nil {
		return
	}

	p.buffer = append(p.buffer, buffer[:n]...)

	if len(p.buffer) < 8 {
		return
	}

	if p.buffer[0] != p.header {
		err = errors.New("invalid header")
		return
	}

	payloadLength := binary.BigEndian.Uint32(p.buffer[2:6])
	if uint32(len(p.buffer)) < payloadLength+6 {
		return
	}

	command = p.buffer[1]
	payload = p.buffer[6 : 6+payloadLength]

	// Trim the buffer
	p.buffer = p.buffer[6+payloadLength:]

	return
}

func New(conn net.Conn, readBufferSize uint, header uint8) *Protocol {
	return &Protocol{
		connection:     conn,
		readBufferSize: readBufferSize,
		header:         header,
		buffer:         make([]byte, 0),
	}
}
