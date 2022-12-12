package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// PendingSellKeyPrefix is the prefix to retrieve all PendingSell
	PendingSellKeyPrefix = "PendingSell/value/"
)

// PendingSellKey returns the store key to retrieve a PendingSell from the index fields
func PendingSellKey(
	name string,
) []byte {
	var key []byte

	nameBytes := []byte(name)
	key = append(key, nameBytes...)
	key = append(key, []byte("/")...)

	return key
}
