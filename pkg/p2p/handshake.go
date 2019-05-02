package p2p

import (
	"encoding/binary"
	"errors"
	"net"
)

type Role int

const (
	joinToken = "some token"
)

// Send token and wait for generated id
func joinNetwork(connection net.Conn) (uint16, uint16, error) {
	p := Proto{connection: connection, buffer: []byte{}}
	// Send the join token
	if err := p.Write(join, []byte(joinToken)); err != nil {
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
func acceptNode(connection net.Conn) (uint16, error) {
	p := Proto{connection: connection, buffer: []byte{}}
	command, data, err := p.Read()
	if err != nil {
		return 0, err
	}

	if command != join {
		return 0, errors.New("no join message received")
	}

	if string(data) != joinToken {
		err := p.Write(rejected, []byte("invalid join token"))
		if err != nil {
			return 0, err
		}
	} else {
		err := p.Write(accepted, []byte{0, 22, 0, 33})
		if err != nil {
			return 0, err
		}
		return 33, nil
	}

	return 0, nil
}
