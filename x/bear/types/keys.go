package types

const (
	// ModuleName defines the module name
	ModuleName = "bear"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_bear"
)

var (
	ParamsKey = []byte("p_bear")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
