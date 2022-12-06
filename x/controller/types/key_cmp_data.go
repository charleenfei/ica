package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// CmpDataKeyPrefix is the prefix to retrieve all CmpData
	CmpDataKeyPrefix = "CmpData/value/"
)

// CmpDataKey returns the store key to retrieve a CmpData from the index fields
func CmpDataKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
