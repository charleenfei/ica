package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// CmpControllerRequestKeyPrefix is the prefix to retrieve all CmpControllerRequest
	CmpControllerRequestKeyPrefix = "CmpControllerRequest/value/"
)

// CmpControllerRequestKey returns the store key to retrieve a CmpControllerRequest from the index fields
func CmpControllerRequestKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
