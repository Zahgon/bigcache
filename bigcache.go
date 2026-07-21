package bigcache

import (
	"context"
)

const (
	minimumEntriesInShard = 10
)

type BigCache struct {
	shards     []*cacheShard
	lifeWindow uint64
	clock      clock
	hash       Hasher
	config     Config
	shardMask  uint64
	close      chan struct{}
}

type Response struct {
	EntryStatus RemoveReason
}

type RemoveReason uint32

const (
	_ RemoveReason = iota

	Expired

	NoSpace

	Deleted
)

func New(ctx context.Context, config Config) (*BigCache, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewBigCache(config Config) (*BigCache, error) { _ = "STUB: not implemented"; return nil, nil }

func newBigCache(ctx context.Context, config Config, clock clock) (*BigCache, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:staticcheck // keep for backward compatibility

func (c *BigCache) Close() error { _ = "STUB: not implemented"; return nil }

func (c *BigCache) Get(key string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *BigCache) GetWithInfo(key string) ([]byte, Response, error) {
	_ = "STUB: not implemented"
	return nil, *new(Response), nil
}

func (c *BigCache) Set(key string, entry []byte) error { _ = "STUB: not implemented"; return nil }

func (c *BigCache) Append(key string, entry []byte) error { _ = "STUB: not implemented"; return nil }

func (c *BigCache) Delete(key string) error { _ = "STUB: not implemented"; return nil }

func (c *BigCache) Reset() error { _ = "STUB: not implemented"; return nil }

func (c *BigCache) ResetStats() error { _ = "STUB: not implemented"; return nil }

func (c *BigCache) Len() int { _ = "STUB: not implemented"; return 0 }

func (c *BigCache) Capacity() int { _ = "STUB: not implemented"; return 0 }

func (c *BigCache) Stats() Stats { _ = "STUB: not implemented"; return *new(Stats) }

func (c *BigCache) KeyMetadata(key string) Metadata {
	_ = "STUB: not implemented"
	return *new(Metadata)
}

func (c *BigCache) Iterator() *EntryInfoIterator { _ = "STUB: not implemented"; return nil }

func (c *BigCache) onEvict(oldestEntry []byte, currentTimestamp uint64, evict func(reason RemoveReason) error) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *BigCache) cleanUp(currentTimestamp uint64) { _ = "STUB: not implemented"; return }

func (c *BigCache) getShard(hashedKey uint64) (shard *cacheShard) {
	_ = "STUB: not implemented"
	return nil
}

func (c *BigCache) providedOnRemove(wrappedEntry []byte, reason RemoveReason) {
	_ = "STUB: not implemented"
	return
}

func (c *BigCache) providedOnRemoveWithReason(wrappedEntry []byte, reason RemoveReason) {
	_ = "STUB: not implemented"
	return
}

func (c *BigCache) notProvidedOnRemove(wrappedEntry []byte, reason RemoveReason) {
	_ = "STUB: not implemented"
	return
}

func (c *BigCache) providedOnRemoveWithMetadata(wrappedEntry []byte, reason RemoveReason) {
	_ = "STUB: not implemented"
	return
}
