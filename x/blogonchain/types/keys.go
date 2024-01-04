package types

const (
	// ModuleName defines the module name
	ModuleName = "blogonchain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_blogonchain"
	PostKey = "Post/value/"
	PostCountKey = "Post/count/"
)

var (
	ParamsKey = []byte("p_blogonchain")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
