package types

const (
	// ModuleName defines the module name
	ModuleName = "ibcchat"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_capability"

	// Version defines the current version the IBC module supports
	Version = "ibcchat-1"

	// PortID is the default port id that module binds to
	PortID = "ibcchat"
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("ibcchat-port-")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	MessageKey      = "Message-value-"
	MessageCountKey = "Message-count-"
)
