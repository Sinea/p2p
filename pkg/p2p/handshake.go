package p2p

import (
	"encoding/binary"
	"fmt"
	"net"
)

type Role int

const (
	Candidate Role = 1 << iota
	Gateway
)

func (s *swarm) handshake(connection net.Conn, role Role) {
	closer := func(c net.Conn) {
		if err := c.Close(); err != nil {
			fmt.Printf("Error closing connection: %s\n", err)
		}
	}
	fmt.Printf("Running handshake with %s\n", connection.RemoteAddr())
	// Write magic header
	magic := make([]byte, 9)
	magic[0] = header

	binary.BigEndian.PutUint64(magic[1:], uint64(s.localID))
	if n, err := connection.Write(magic); err != nil {
		fmt.Printf("Error writing magic: %s\n", err)
		closer(connection)
		return
	} else if n < len(magic) {
		fmt.Printf("Wrote only %d bytes\n", n)
		closer(connection)
		return
	}

	receivedMagic := make([]byte, 9)

	if n, err := connection.Read(receivedMagic); err != nil {
		fmt.Printf("Error reading magic: %s\n", err)
		closer(connection)
		return
	} else if n != 9 {
		fmt.Printf("Read only %d bytes\n", n)
		closer(connection)
		return
	}

	fmt.Printf("Received magic %d\n", receivedMagic)

	if receivedMagic[0] != header {
		fmt.Printf("Received %d as magic header :(. Bye!\n", receivedMagic[0])
		return
	}

	remote := PeerID(binary.BigEndian.Uint64(receivedMagic[1:]))
	fmt.Printf("Hello %d\n", remote)

	pp := &peer{remote, connection, make([]byte, 0)}
	s.peers[remote] = pp
	s.nodes[remote] = pp
}
