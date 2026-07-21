package bigcache

func newDefaultHasher() Hasher { _ = "STUB: not implemented"; return *new(Hasher) }

type fnv64a struct{}

const (
	offset64 = 14695981039346656037

	prime64 = 1099511628211
)

func (f fnv64a) Sum64(key string) uint64 { _ = "STUB: not implemented"; return 0 }
