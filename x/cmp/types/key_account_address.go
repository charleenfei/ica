package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// AccountAddressKeyPrefix is the prefix to retrieve all AccountAddress
	AccountAddressKeyPrefix = "AccountAddress/value/"
)

// AccountAddressKey returns the store key to retrieve a AccountAddress from the index fields
func AccountAddressKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
