package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// CmpHostResultKeyPrefix is the prefix to retrieve all CmpHostResult
	CmpHostResultKeyPrefix = "CmpHostResult/value/"
)

// CmpHostResultKey returns the store key to retrieve a CmpHostResult from the index fields
func CmpHostResultKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
