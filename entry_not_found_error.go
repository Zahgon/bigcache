package bigcache

import "errors"

var (
	ErrEntryNotFound = errors.New("Entry not found") //nolint:staticcheck // keep for backward compatibility
)
