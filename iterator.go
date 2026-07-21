package bigcache

import (
	"sync"
)

type iteratorError string

func (e iteratorError) Error() string { _ = "STUB: not implemented"; return "" }

const ErrInvalidIteratorState = iteratorError("Iterator is in invalid state. Use SetNext() to move to next position")

const ErrCannotRetrieveEntry = iteratorError("Could not retrieve entry from cache")

var emptyEntryInfo = EntryInfo{}

type EntryInfo struct {
	timestamp uint64
	hash      uint64
	key       string
	value     []byte
	err       error
}

func (e EntryInfo) Key() string { _ = "STUB: not implemented"; return "" }

func (e EntryInfo) Hash() uint64 { _ = "STUB: not implemented"; return 0 }

func (e EntryInfo) Timestamp() uint64 { _ = "STUB: not implemented"; return 0 }

func (e EntryInfo) Value() []byte { _ = "STUB: not implemented"; return nil }

type EntryInfoIterator struct {
	mutex            sync.Mutex
	cache            *BigCache
	currentShard     int
	currentIndex     int
	currentEntryInfo EntryInfo
	elements         []uint64
	elementsCount    int
	valid            bool
}

func (it *EntryInfoIterator) SetNext() bool { _ = "STUB: not implemented"; return false }

func (it *EntryInfoIterator) setCurrentEntry() bool { _ = "STUB: not implemented"; return false }

func newIterator(cache *BigCache) *EntryInfoIterator { _ = "STUB: not implemented"; return nil }

func (it *EntryInfoIterator) Value() (EntryInfo, error) {
	_ = "STUB: not implemented"
	return *new(EntryInfo), nil
}
