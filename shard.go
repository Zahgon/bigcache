package bigcache

import (
	"sync"

	"github.com/allegro/bigcache/v3/queue"
)

type onRemoveCallback func(wrappedEntry []byte, reason RemoveReason)

type Metadata struct {
	RequestCount uint32
}

type cacheShard struct {
	hashmap     map[uint64]uint64
	entries     queue.BytesQueue
	lock        sync.RWMutex
	entryBuffer []byte
	onRemove    onRemoveCallback

	isVerbose    bool
	statsEnabled bool
	logger       Logger
	clock        clock
	lifeWindow   uint64

	hashmapStats map[uint64]uint32
	stats        Stats
	cleanEnabled bool
}

func (s *cacheShard) getWithInfo(key string, hashedKey uint64) (entry []byte, resp Response, err error) {
	_ = "STUB: not implemented"
	return nil, *new(Response), nil
}

func (s *cacheShard) get(key string, hashedKey uint64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *cacheShard) getWrappedEntry(hashedKey uint64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *cacheShard) getValidWrapEntry(key string, hashedKey uint64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *cacheShard) set(key string, hashedKey uint64, entry []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *cacheShard) addNewWithoutLock(key string, hashedKey uint64, entry []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *cacheShard) setWrappedEntryWithoutLock(currentTimestamp uint64, w []byte, hashedKey uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *cacheShard) append(key string, hashedKey uint64, entry []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *cacheShard) del(hashedKey uint64) error { _ = "STUB: not implemented"; return nil }

func (s *cacheShard) onEvict(oldestEntry []byte, currentTimestamp uint64, evict func(reason RemoveReason) error) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *cacheShard) isExpired(oldestEntry []byte, currentTimestamp uint64) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *cacheShard) cleanUp(currentTimestamp uint64) { _ = "STUB: not implemented"; return }

func (s *cacheShard) getEntry(hashedKey uint64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *cacheShard) copyHashedKeys() (keys []uint64, next int) {
	_ = "STUB: not implemented"
	return nil, 0
}

func (s *cacheShard) removeOldestEntry(reason RemoveReason) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *cacheShard) reset(config Config) { _ = "STUB: not implemented"; return }

func (s *cacheShard) resetStats() { _ = "STUB: not implemented"; return }

func (s *cacheShard) len() int { _ = "STUB: not implemented"; return 0 }

func (s *cacheShard) capacity() int { _ = "STUB: not implemented"; return 0 }

func (s *cacheShard) getStats() Stats { _ = "STUB: not implemented"; return *new(Stats) }

func (s *cacheShard) getKeyMetadataWithLock(key uint64) Metadata {
	_ = "STUB: not implemented"
	return *new(Metadata)
}

func (s *cacheShard) getKeyMetadata(key uint64) Metadata {
	_ = "STUB: not implemented"
	return *new(Metadata)
}

func (s *cacheShard) hit(key uint64) { _ = "STUB: not implemented"; return }

func (s *cacheShard) hitWithoutLock(key uint64) { _ = "STUB: not implemented"; return }

func (s *cacheShard) miss() { _ = "STUB: not implemented"; return }

func (s *cacheShard) delhit() { _ = "STUB: not implemented"; return }

func (s *cacheShard) delmiss() { _ = "STUB: not implemented"; return }

func (s *cacheShard) collision() { _ = "STUB: not implemented"; return }

func initNewShard(config Config, callback onRemoveCallback, clock clock) *cacheShard {
	_ = "STUB: not implemented"
	return nil
}
