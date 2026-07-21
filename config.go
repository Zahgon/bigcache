package bigcache

import "time"

type Config struct {
	Shards int

	LifeWindow time.Duration

	CleanWindow time.Duration

	MaxEntriesInWindow int

	MaxEntrySize int

	StatsEnabled bool

	Verbose bool

	Hasher Hasher

	HardMaxCacheSize int

	OnRemove func(key string, entry []byte)

	OnRemoveWithMetadata func(key string, entry []byte, keyMetadata Metadata)

	OnRemoveWithReason func(key string, entry []byte, reason RemoveReason)

	onRemoveFilter int

	Logger Logger
}

func DefaultConfig(eviction time.Duration) Config { _ = "STUB: not implemented"; return *new(Config) }

func (c Config) initialShardSize() int { _ = "STUB: not implemented"; return 0 }

func (c Config) maximumShardSizeInBytes() int { _ = "STUB: not implemented"; return 0 }

func (c Config) OnRemoveFilterSet(reasons ...RemoveReason) Config {
	_ = "STUB: not implemented"
	return *new(Config)
}
