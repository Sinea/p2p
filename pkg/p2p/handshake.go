package p2p

import (
	"bytes"
	"encoding/binary"
	"errors"
	"net"
)

// Send token and wait for generated id
func joinNetwork(connection net.Conn, token []byte) (uint16, uint16, error) {
	p := NewProtocol(connection)
	// Send the join token
	if err := p.Write(join, token); err != nil {
		return 0, 0, err
	}

	// Wait for a reply
	command, data, err := p.Read()
	if err != nil {
		return 0, 0, err
	}

	if command == accepted {
		myID := binary.BigEndian.Uint16(data[0:2])
		remoteID := binary.BigEndian.Uint16(data[2:4])
		return myID, remoteID, nil
	}

	return 0, 0, errors.New("error joining network")
}

// Wait for token, check the token, send back a generated id
func acceptNode(connection net.Conn, token []byte) (uint16, error) {
	p := NewProtocol(connection)
	command, receivedToken, err := p.Read()
	if err != nil {
		return 0, err
	}

	if command != join {
		return 0, errors.New("no join message received")
	}

	if bytes.Compare(token, receivedToken) != 0 {
		err := p.Write(rejected, []byte("invalid join token"))
		if err != nil {
			return 0, err
		}
		return 0, errors.New("invalid join token")
	} else {
		err := p.Write(accepted, []byte{0, 33, 0, 22})
		if err != nil {
			return 0, err
		}
		return 33, nil
	}
}
