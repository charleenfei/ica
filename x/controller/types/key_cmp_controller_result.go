package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// CmpControllerResultKeyPrefix is the prefix to retrieve all CmpControllerResult
	CmpControllerResultKeyPrefix = "CmpControllerResult/value/"
)

// CmpControllerResultKey returns the store key to retrieve a CmpControllerResult from the index fields
func CmpControllerResultKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
