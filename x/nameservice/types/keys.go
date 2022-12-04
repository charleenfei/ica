package types

const (
	// ModuleName defines the module name
	ModuleName = "nameservice"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_nameservice"

	// For event emitting
	CmpHostRequestEventType    = "cmp-host-request" // Indicates what event type to listen to
	CmpHostRequestEventCreator = "creator"          // Subsidiary information
	CmpHostRequestEventId      = "request-id"       // which request if multiple requests happened
	CmpHostRequestItem         = "request-item"     // Is it relevant to me?
	CmpHostRequestBid          = "request-bid"      // Is it relevant to me?

)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	TestminKey = "Testmin/value/"
)
