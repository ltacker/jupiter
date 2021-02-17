package types

const (
	// ModuleName defines the module name
	ModuleName = "laugh"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_capability"

	// Version defines the current version the IBC module supports
	Version = "laugh-1"

	// PortID is the default port id that module binds to
	PortID = "laugh"
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("laugh-port-")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	HahasentKey      = "Hahasent-value-"
	HahasentCountKey = "Hahasent-count-"
)

const (
	HihisentKey      = "Hihisent-value-"
	HihisentCountKey = "Hihisent-count-"
)

const (
	HohosentKey      = "Hohosent-value-"
	HohosentCountKey = "Hohosent-count-"
)

const (
	HahaKey      = "Haha-value-"
	HahaCountKey = "Haha-count-"
)

const (
	HihiKey      = "Hihi-value-"
	HihiCountKey = "Hihi-count-"
)

const (
	HohoKey      = "Hoho-value-"
	HohoCountKey = "Hoho-count-"
)
