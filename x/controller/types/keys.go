package types

const (
	// ModuleName defines the module name
	ModuleName = "controller"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_controller"

	// For event emitting
	CmpControllerRequestEventType    = "cmp-controller-request" // Indicates what event type to listen to
	CmpControllerRequestEventCreator = "request-creator"        // Subsidiary information
	CmpControllerRequestEventId      = "request-id"             // which request if multiple requests happened
	CmpControllerRequestMetaData     = "request-metadata"       // Extra data for other purpose

	// for event querying result of cmp request
	CmpResultRequestEventType = "cmp-result-request"
	CmpResultRequestId        = "request-id"
	CmpResultRequestData      = "request-data"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
