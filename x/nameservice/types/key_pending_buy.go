package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// PendingBuyKeyPrefix is the prefix to retrieve all PendingBuy
	PendingBuyKeyPrefix = "PendingBuy/value/"
)

// PendingBuyKey returns the store key to retrieve a PendingBuy from the index fields
func PendingBuyKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
