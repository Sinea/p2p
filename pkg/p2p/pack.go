package p2p

import (
	"encoding/binary"
)

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
